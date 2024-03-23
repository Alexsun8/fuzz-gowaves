package fuzz_WE

import "C"
import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"

	_ "github.com/dvyukov/go-fuzz/go-fuzz-dep"
	_ "github.com/google/gofuzz"
	"github.com/wavesplatform/gowaves/pkg/api"
	"github.com/wavesplatform/gowaves/pkg/node_init"
)

var sharedApi *api.NodeApi
var standardErrors = []string{
	"looking for beginning of value",
	"unexpected end of JSON input",
	"looking for beginning of object key string",
	"json: cannot unmarshal",
	"invalid character",
	"unknown transaction type",
}

func containsStandardError(err error) bool {
	for _, se := range standardErrors {
		if strings.Contains(err.Error(), se) {
			return true
		}
	}
	return false
}

func NodeInitForFuzz() bool {
	// Запуск сервера или другие инициализации, если они нужны
	wg := new(sync.WaitGroup)
	defer wg.Wait()
	wg.Add(1)
	//ch := make(chan struct{}, 1)
	ch := make(
		chan node_init.ChanMessage,
		1,
	)
	go func() {
		defer wg.Done()
		node_init.RunNodeInit(ch)
	}()
	msg := <-ch
	fmt.Sprintln("tests's started")
	sharedApi = msg.NodeAPI
	return true
}

// func FuzzTransactionsBroadcast(data []byte) []byte {
// func FuzzTransactionsBroadcast(data []byte) int {
func FuzzTransactionsBroadcast(data []byte) int {
	// Инициализация HTTP Recorder
	w := httptest.NewRecorder()

	// Преобразование данных JSON в строку
	fuzzData := string(data) // (!) добавить структурированность, если не json, вернуть -1)

	_, err := json.Marshal(fuzzData)

	if err != nil {
		fmt.Println("error:", err)
		return -1
	}

	// Создание нового запроса POST
	req := httptest.NewRequest("POST", "/transactions/broadcast", bytes.NewBuffer([]byte(fuzzData)))

	// Вызов функции TransactionsBroadcast и обработка ошибок
	if err := sharedApi.TransactionsBroadcast(w, req); err != nil {
		// Если возникла ошибка, сообщите о ней
		if containsStandardError(err) {
			return -1
		}

		fE := fmt.Sprintf("Error calling TransactionsBroadcast: %v\n", err) // ToDo: ошибки "unexpected end of JSON input " и "invalid character '\x01" должны игнориться
		fmt.Printf(fE)
		return 1
	}
	// Проверка кода статуса ответа (пример: 200 - OK)
	if w.Code != http.StatusOK {
		//t.Errorf("Expected status 200, got %d, body: %s", w.Code, w.Body)
		sE := fmt.Sprintf("Expected status 200, got %d, body: %s", w.Code, w.Body)
		fmt.Printf(sE)

		return 1
	}
	fmt.Printf("Answer: %d", w.Code)

	// Возврат кода статуса ответа HTTP
	return 0
}
