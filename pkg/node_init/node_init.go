package node_init

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net/http"
	_ "net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/wavesplatform/gowaves/pkg/api"
	"github.com/wavesplatform/gowaves/pkg/crypto"
	"github.com/wavesplatform/gowaves/pkg/grpc/server"
	"github.com/wavesplatform/gowaves/pkg/libs/microblock_cache"
	"github.com/wavesplatform/gowaves/pkg/libs/ntptime"
	"github.com/wavesplatform/gowaves/pkg/libs/runner"
	"github.com/wavesplatform/gowaves/pkg/logging"
	"github.com/wavesplatform/gowaves/pkg/metrics"
	"github.com/wavesplatform/gowaves/pkg/miner"
	"github.com/wavesplatform/gowaves/pkg/miner/scheduler"
	"github.com/wavesplatform/gowaves/pkg/miner/utxpool"
	"github.com/wavesplatform/gowaves/pkg/node"
	"github.com/wavesplatform/gowaves/pkg/node/blocks_applier"
	"github.com/wavesplatform/gowaves/pkg/node/messages"
	"github.com/wavesplatform/gowaves/pkg/node/network"
	"github.com/wavesplatform/gowaves/pkg/node/peers"
	peersPersistentStorage "github.com/wavesplatform/gowaves/pkg/node/peers/storage"
	"github.com/wavesplatform/gowaves/pkg/p2p/peer"
	"github.com/wavesplatform/gowaves/pkg/proto"
	"github.com/wavesplatform/gowaves/pkg/services"
	"github.com/wavesplatform/gowaves/pkg/settings"
	"github.com/wavesplatform/gowaves/pkg/state"
	"github.com/wavesplatform/gowaves/pkg/types"
	"github.com/wavesplatform/gowaves/pkg/util/common"
	"github.com/wavesplatform/gowaves/pkg/util/fdlimit"
	"github.com/wavesplatform/gowaves/pkg/versioning"
	"github.com/wavesplatform/gowaves/pkg/wallet"
)

const (
	mb             = 1 << 20
	defaultTimeout = 30 * time.Second
	reserve        = 10
)

var defaultPeers = map[string]string{
	"mainnet":  "34.253.153.4:6868,168.119.116.189:6868,135.181.87.72:6868,162.55.39.115:6868,168.119.155.201:6868",
	"testnet":  "159.69.126.149:6868,94.130.105.239:6868,159.69.126.153:6868,94.130.172.201:6868,35.157.247.122:6868",
	"stagenet": "88.99.185.128:6868,49.12.15.166:6868,95.216.205.3:6868,88.198.179.16:6868,52.58.254.101:6868",
}

type config struct {
	logLevel                   zapcore.Level
	logDevelopment             bool
	logNetwork                 bool
	logNetworkData             bool
	logFSM                     bool
	statePath                  string
	blockchainType             string
	peerAddresses              string
	declAddr                   string
	nodeName                   string
	cfgPath                    string
	apiAddr                    string
	apiKey                     string
	apiMaxConnections          int
	rateLimiterOptions         string
	grpcAddr                   string
	grpcAPIMaxConnections      int
	enableMetaMaskAPI          bool
	enableMetaMaskAPILog       bool
	enableGrpcAPI              bool
	blackListResidenceTime     time.Duration
	buildExtendedAPI           bool
	serveExtendedAPI           bool
	buildStateHashes           bool
	bindAddress                string
	disableOutgoingConnections bool
	minerVoteFeatures          string
	disableBloomFilter         bool
	reward                     string
	obsolescencePeriod         time.Duration
	walletPath                 string
	walletPassword             string
	limitAllConnections        uint
	minPeersMining             int
	disableMiner               bool
	profiler                   bool
	prometheus                 string
	metricsID                  int
	metricsURL                 string
	dropPeers                  bool
	dbFileDescriptors          int
	newConnectionsLimit        int
	disableNTP                 bool
	microblockInterval         time.Duration
}

func (c *config) logParameters() {
	zap.S().Debugf("log-level: %s", c.logLevel)
	zap.S().Debugf("log-dev: %t", c.logDevelopment)
	zap.S().Debugf("log-network: %t", c.logNetwork)
	zap.S().Debugf("log-fsm: %t", c.logFSM)
	zap.S().Debugf("state-path: %s", c.statePath)
	zap.S().Debugf("blockchain-type: %s", c.blockchainType)
	zap.S().Debugf("peers: %s", c.peerAddresses)
	zap.S().Debugf("declared-address: %s", c.declAddr)
	zap.S().Debugf("api-address: %s", c.apiAddr)
	zap.S().Debugf("api-key: %s", crypto.MustKeccak256([]byte(c.apiKey)).Hex())
	zap.S().Debugf("grpc-address: %s", c.grpcAddr)
	zap.S().Debugf("enable-grpc-api: %t", c.enableGrpcAPI)
	zap.S().Debugf("black-list-residence-time: %s", c.blackListResidenceTime)
	zap.S().Debugf("build-extended-api: %t", c.buildExtendedAPI)
	zap.S().Debugf("serve-extended-api: %t", c.serveExtendedAPI)
	zap.S().Debugf("build-state-hashes: %t", c.buildStateHashes)
	zap.S().Debugf("bind-address: %s", c.bindAddress)
	zap.S().Debugf("vote: %s", c.minerVoteFeatures)
	zap.S().Debugf("reward: %s", c.reward)
	zap.S().Debugf("obsolescence: %s", c.obsolescencePeriod)
	zap.S().Debugf("disable-miner %t", c.disableMiner)
	zap.S().Debugf("wallet-path: %s", c.walletPath)
	zap.S().Debugf("hashed wallet-password: %s", crypto.MustKeccak256([]byte(c.walletPassword)).Hex())
	zap.S().Debugf("limit-connections: %d", c.limitAllConnections)
	zap.S().Debugf("profiler: %t", c.profiler)
	zap.S().Debugf("disable-bloom: %t", c.disableBloomFilter)
	zap.S().Debugf("drop-peers: %t", c.dropPeers)
	zap.S().Debugf("db-file-descriptors: %v", c.dbFileDescriptors)
	zap.S().Debugf("new-connections-limit: %v", c.newConnectionsLimit)
	zap.S().Debugf("enable-metamask: %t", c.enableMetaMaskAPI)
	zap.S().Debugf("disable-ntp: %t", c.disableNTP)
	zap.S().Debugf("microblock-interval: %s", c.microblockInterval)
}

// -blockchain-type stagenet -state-path /Users/alexsun/BMSTU/diplom/stagenet
// -api-address 127.0.0.1:8083 -no-connections -peers 127.0.0.1:6868 -build-extended-api
// -build-state-hashes -serve-extended-api -api-key="key"

func (c *config) hardcode_config() {
	const (
		defaultBlacklistResidenceDuration = 5 * time.Minute
		defaultObsolescenceDuration       = 4 * time.Hour
		defaultConnectionsLimit           = 10000
		defaultNewConnectionLimit         = 10000
		defaultMicroblockInterval         = 1 * time.Second //! (!)
	)
	c.logLevel = zapcore.InfoLevel
	c.logDevelopment = false
	c.logNetwork = false
	c.logNetworkData = false
	c.logFSM = false
	c.statePath = "/com.docker.devenvironments.code/stagenet"
	c.blockchainType = "stagenet"
	c.peerAddresses = "127.0.0.1:6868"
	c.declAddr = ""
	c.nodeName = "fuzz-gowaves"
	c.cfgPath = ""
	c.apiAddr = "127.0.0.1:8083"
	c.apiKey = "key"
	c.apiMaxConnections = api.DefaultMaxConnections
	c.rateLimiterOptions = ""
	c.grpcAddr = "127.0.0.1:7475" // "Address for gRPC API."
	c.grpcAPIMaxConnections = server.DefaultMaxConnections
	c.enableMetaMaskAPI = true
	c.enableMetaMaskAPILog = false
	c.enableGrpcAPI = false
	c.blackListResidenceTime = defaultBlacklistResidenceDuration
	c.buildExtendedAPI = true
	c.serveExtendedAPI = true
	c.buildStateHashes = true
	c.bindAddress = ""
	c.disableOutgoingConnections = true
	c.minerVoteFeatures = ""
	c.disableBloomFilter = false
	c.reward = ""
	c.obsolescencePeriod = defaultObsolescenceDuration
	c.walletPath = ""
	c.walletPassword = ""
	c.limitAllConnections = defaultConnectionsLimit
	c.minPeersMining = 1
	c.disableMiner = false
	c.profiler = false
	c.prometheus = ""
	c.metricsID = -1
	c.metricsURL = ""
	c.dropPeers = false
	c.dbFileDescriptors = state.DefaultOpenFilesCacheCapacity
	c.newConnectionsLimit = defaultNewConnectionLimit
	c.disableNTP = false
	c.microblockInterval = defaultMicroblockInterval

}

type Scheduler interface {
	Mine() chan scheduler.Emit
	types.Scheduler
	Emits() []scheduler.Emit
}

type ChanMessage struct {
	//Ctx     context.Context
	NodeAPI *api.NodeApi
	// Добавьте дополнительные поля, если необходимо
}

// func run(ch chan<- struct{}) {
func RunNodeInit(ch chan<- ChanMessage) {
	nc := new(config)
	nc.hardcode_config() // nc.parse() => nc.hardcode_config, for deleting cmd-flags
	logger := logging.SetupLogger(nc.logLevel,
		logging.DevelopmentFlag(nc.logDevelopment),
		logging.NetworkFilter(nc.logNetwork),
		logging.NetworkDataFilter(nc.logNetworkData),
		logging.FSMFilter(nc.logFSM),
	)
	defer func() {
		err := logger.Sync()
		if err != nil && errors.Is(err, os.ErrInvalid) {
			panic(fmt.Sprintf("Failed to close logging subsystem: %v\n", err))
		}
	}()

	zap.S().Infof("Gowaves Node version: %s", versioning.Version)

	maxFDs, err := fdlimit.MaxFDs()
	if err != nil {
		zap.S().Fatalf("Initialization failure: %v", err)
	}
	_, err = fdlimit.RaiseMaxFDs(maxFDs)
	if err != nil {
		zap.S().Fatalf("Initialization failure: %v", err)
	}
	if m := int(maxFDs) - int(nc.limitAllConnections) - reserve; nc.dbFileDescriptors > m {
		zap.S().Fatalf(
			"Invalid 'db-file-descriptors' flag value (%d). Value shall be less or equal to %d.",
			nc.dbFileDescriptors, m)
	}

	if nc.profiler {
		zap.S().Infof("Starting built-in profiler on 'http://localhost:6060/debug/pprof/'")
		go func() {
			pprofMux := http.NewServeMux()
			// taken from "net/http/pprof" init()
			pprofMux.HandleFunc("/debug/pprof/", pprof.Index)
			pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
			pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)
			s := &http.Server{
				Addr:              "localhost:6060",
				Handler:           pprofMux,
				ReadHeaderTimeout: defaultTimeout,
				ReadTimeout:       defaultTimeout,
			}
			zap.S().Warn(s.ListenAndServe())
		}()
	}

	ctx, done := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer done()

	if nc.metricsURL != "" && nc.metricsID != -1 {
		err = metrics.Start(ctx, nc.metricsID, nc.metricsURL)
		if err != nil {
			zap.S().Warnf("Metrics reporting failed to start: %v", err)
			zap.S().Warn("Proceeding without reporting any metrics")
		} else {
			zap.S().Info("Metrics reporting activated")
		}
	}

	nc.logParameters()

	var cfg *settings.BlockchainSettings
	if nc.cfgPath != "" {
		var f *os.File
		f, err = os.Open(nc.cfgPath)
		if err != nil {
			zap.S().Fatalf("Failed to open configuration file: %v", err)
		}
		defer func() { _ = f.Close() }()
		cfg, err = settings.ReadBlockchainSettings(f)
		if err != nil {
			zap.S().Fatalf("Failed to read configuration file: %v", err)
		}
	} else {
		cfg, err = settings.BlockchainSettingsByTypeName(nc.blockchainType)
		if err != nil {
			zap.S().Errorf("Failed to get blockchain settings: %v", err)
			return
		}
	}

	conf := &settings.NodeSettings{}
	err = settings.ApplySettings(conf, FromArgs(cfg.AddressSchemeCharacter, nc), settings.FromJavaEnviron)
	if err != nil {
		zap.S().Errorf("Failed to apply node settings: %v", err)
		return
	}

	err = conf.Validate()
	if err != nil {
		zap.S().Errorf("Failed to validate node settings: %v", err)
		return
	}

	var wal types.EmbeddedWallet = wallet.NewEmbeddedWallet(wallet.NewLoader(nc.walletPath),
		wallet.NewWallet(), cfg.AddressSchemeCharacter)
	if nc.walletPassword != "" {
		if err = wal.Load([]byte(nc.walletPassword)); err != nil {
			zap.S().Errorf("Failed to load wallet: %v", err)
			return
		}
	}

	path := nc.statePath
	if path == "" {
		path, err = common.GetStatePath()
		if err != nil {
			zap.S().Errorf("Failed to get state path: %v", err)
			return
		}
	}

	reward, err := miner.ParseReward(nc.reward)
	if err != nil {
		zap.S().Errorf("Failed to parse '-reward': %v", err)
		return
	}

	ntpTime, err := getNtp(ctx, nc.disableNTP)
	if err != nil {
		zap.S().Errorf("Failed to get NTP time: %v", err)
		return
	}

	params := state.DefaultStateParams()
	params.StorageParams.DbParams.OpenFilesCacheCapacity = nc.dbFileDescriptors
	params.StoreExtendedApiData = nc.buildExtendedAPI
	params.ProvideExtendedApi = nc.serveExtendedAPI
	params.BuildStateHashes = nc.buildStateHashes
	params.Time = ntpTime
	params.DbParams.BloomFilterParams.Disable = nc.disableBloomFilter

	st, err := state.NewState(path, true, params, cfg)
	if err != nil {
		zap.S().Error("Failed to initialize node's state: %v", err)
		return
	}

	features, err := miner.ParseVoteFeatures(nc.minerVoteFeatures)
	if err != nil {
		zap.S().Errorf("Failed to parse '-vote': %v", err)
		return
	}

	features, err = miner.ValidateFeatures(st, features)
	if err != nil {
		zap.S().Errorf("Failed to validate features: %v", err)
		return
	}

	// Check if we need to start serving extended API right now.
	if err := node.MaybeEnableExtendedApi(st, ntpTime); err != nil {
		zap.S().Errorf("Failed to enable extended API: %v", err)
		return
	}

	async := runner.NewAsync()
	logRunner := runner.NewLogRunner(async)

	declAddr := proto.NewTCPAddrFromString(conf.DeclaredAddr)
	bindAddr := proto.NewTCPAddrFromString(nc.bindAddress)

	utxValidator, err := utxpool.NewValidator(st, ntpTime, nc.obsolescencePeriod)
	if err != nil {
		zap.S().Errorf("Failed to initialize UTX: %v", err)
		return
	}
	utx := utxpool.New(uint64(1024*mb), utxValidator, cfg)
	parent := peer.NewParent()

	nodeNonce, err := rand.Int(rand.Reader, new(big.Int).SetUint64(math.MaxInt32))
	if err != nil {
		zap.S().Errorf("Failed to get node's nonce: %v", err)
		return
	}
	peerSpawnerImpl := peers.NewPeerSpawner(parent, conf.WavesNetwork, declAddr, nc.nodeName,
		nodeNonce.Uint64(), proto.ProtocolVersion)
	peerStorage, err := peersPersistentStorage.NewCBORStorage(nc.statePath, time.Now())
	if err != nil {
		zap.S().Errorf("Failed to open or create peers storage: %v", err)
		return
	}
	if nc.dropPeers {
		if err := peerStorage.DropStorage(); err != nil {
			zap.S().Errorf("Failed to drop peers storage. Drop peers storage manually. Err: %v", err)
			return
		}
		zap.S().Info("Successfully dropped peers storage")
	}

	peerManager := peers.NewPeerManager(
		peerSpawnerImpl,
		peerStorage,
		int(nc.limitAllConnections/2),
		proto.ProtocolVersion,
		conf.WavesNetwork,
		!nc.disableOutgoingConnections,
		nc.newConnectionsLimit,
		nc.blackListResidenceTime,
	)
	go peerManager.Run(ctx)

	var minerScheduler Scheduler
	if nc.disableMiner {
		minerScheduler = scheduler.DisabledScheduler{}
	} else {
		minerScheduler, err = scheduler.NewScheduler(
			st,
			wal,
			cfg,
			ntpTime,
			scheduler.NewMinerConsensus(peerManager, nc.minPeersMining),
			nc.obsolescencePeriod,
		)
		if err != nil {
			zap.S().Errorf("Failed to initialize miner scheduler: %v", err)
			return
		}
	}
	blockApplier := blocks_applier.NewBlocksApplier()

	svs := services.Services{
		State:           st,
		Peers:           peerManager,
		Scheduler:       minerScheduler,
		BlocksApplier:   blockApplier,
		UtxPool:         utx,
		Scheme:          cfg.AddressSchemeCharacter,
		LoggableRunner:  logRunner,
		Time:            ntpTime,
		Wallet:          wal,
		MicroBlockCache: microblock_cache.NewMicroblockCache(),
		InternalChannel: messages.NewInternalChannel(),
		MinPeersMining:  nc.minPeersMining,
		SkipMessageList: parent.SkipMessageList,
	}

	mine := miner.NewMicroblockMiner(svs, features, reward)
	go miner.Run(ctx, mine, minerScheduler, svs.InternalChannel)

	ntw, networkInfoCh := network.NewNetwork(svs, parent, nc.obsolescencePeriod)
	go ntw.Run(ctx)

	n := node.NewNode(svs, declAddr, bindAddr, nc.microblockInterval)
	go n.Run(ctx, parent, svs.InternalChannel, networkInfoCh, ntw.SyncPeer())

	go minerScheduler.Reschedule()

	if len(conf.Addresses) > 0 {
		addresses := strings.Split(conf.Addresses, ",")
		for _, addr := range addresses {
			tcpAddr := proto.NewTCPAddrFromString(addr)
			if tcpAddr.Empty() {
				// That means that configuration parameter is invalid
				zap.S().Errorf("Failed to parse TCPAddr from string %q", tcpAddr.String())
				return
			}
			if err = peerManager.AddAddress(ctx, tcpAddr); err != nil {
				// That means that we have problems with peers storage
				zap.S().Errorf("Failed to add addres into know peers storage: %v", err)
				return
			}
		}
	}

	app, err := api.NewApp(nc.apiKey, minerScheduler, svs)
	if err != nil {
		zap.S().Errorf("Failed to initialize application: %v", err)
		return
	}

	webApi := api.NewNodeApi(app, st, n)
	go func() {
		zap.S().Infof("Starting node HTTP API on '%v'", conf.HttpAddr)
		if err = api.Run(ctx, conf.HttpAddr, webApi, apiRunOptsFromCLIFlags(nc)); err != nil {
			zap.S().Errorf("Failed to start API: %v", err)
		}
	}()

	go func() {
		if nc.prometheus != "" {
			h := http.NewServeMux()
			h.Handle("/metrics", promhttp.Handler())
			s := &http.Server{
				Addr:              nc.prometheus,
				Handler:           h,
				ReadHeaderTimeout: defaultTimeout,
				ReadTimeout:       defaultTimeout,
			}
			zap.S().Infof("Starting node metrics endpoint on '%v'", nc.prometheus)
			_ = s.ListenAndServe()
		}
	}()

	if nc.enableGrpcAPI {
		var srv *server.Server
		srv, err = server.NewServer(svs)
		if err != nil {
			zap.S().Errorf("Failed to create gRPC server: %v", err)
		}
		go func() {
			err = srv.Run(ctx, conf.GrpcAddr, grpcAPIRunOptsFromCLIFlags(nc))
			if err != nil {
				zap.S().Errorf("grpcServer.Run(): %v", err)
			}
		}()
	}
	<-time.After(1 * time.Second)
	zap.S().Info("Ready for tests")
	if ch != nil {
		select {
		//case ch <- struct{}{}:
		case ch <- ChanMessage{webApi}:
		default:
			// no-op
		}
	}
	zap.S().Info("smth went wrong!!!")

	<-ctx.Done()
	zap.S().Info("User termination in progress...")
	n.Close()
	<-time.After(1 * time.Second)
}

func FromArgs(scheme proto.Scheme, c *config) func(s *settings.NodeSettings) error {
	return func(s *settings.NodeSettings) error {
		s.DeclaredAddr = c.declAddr
		s.HttpAddr = c.apiAddr
		s.GrpcAddr = c.grpcAddr
		s.WavesNetwork = proto.NetworkStrFromScheme(scheme)
		s.Addresses = c.peerAddresses
		if c.peerAddresses == "" {
			s.Addresses = defaultPeers[c.blockchainType]
		}
		return nil
	}
}

func apiRunOptsFromCLIFlags(c *config) *api.RunOptions {
	// TODO: add more run flags to CLI flags
	opts := api.DefaultRunOptions()
	opts.MaxConnections = c.apiMaxConnections
	if c.enableMetaMaskAPI {
		if c.buildExtendedAPI {
			opts.EnableMetaMaskAPI = c.enableMetaMaskAPI
			opts.EnableMetaMaskAPILog = c.enableMetaMaskAPILog
		} else {
			zap.S().Warn("'enable-metamask' flag requires activated 'build-extended-api' flag")
		}
	}
	if c.rateLimiterOptions != "" {
		rlo, err := api.NewRateLimiterOptionsFromString(c.rateLimiterOptions)
		if err == nil {
			opts.RateLimiterOpts = rlo
		} else {
			zap.S().Errorf("Invalid rate limiter options '%s': %v", c.rateLimiterOptions, err)
		}
	}
	return opts
}

func grpcAPIRunOptsFromCLIFlags(c *config) *server.RunOptions {
	opts := server.DefaultRunOptions()
	opts.MaxConnections = c.grpcAPIMaxConnections
	return opts
}

func getNtp(ctx context.Context, disable bool) (types.Time, error) {
	if disable {
		return ntptime.Stub{}, nil
	}
	tm, err := ntptime.TryNew("pool.ntp.org", 10)
	if err != nil {
		return nil, err
	}
	go tm.Run(ctx, 2*time.Minute)
	return tm, nil
}
