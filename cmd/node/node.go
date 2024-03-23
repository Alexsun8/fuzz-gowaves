package main

import (
	_ "net/http"

	"github.com/wavesplatform/gowaves/pkg/node_init"
)

func main() {
	node_init.RunNodeInit(nil)
}
