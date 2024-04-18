package fuzz_WE

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// )

// // Пример структуры для данных, передаваемых в API
// type TestData struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// func FuzzBlocksScoreAt(data []byte) int {
// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Преобразование данных JSON в структуру TestData
// 	var testData TestData
// 	if err := json.Unmarshal(data, &testData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	// Создание нового запроса GET
// 	req := httptest.NewRequest("GET", fmt.Sprintf("/go/blocks/score/at/%d", testData.ID), nil)

// 	// Вызов функции API и обработка ошибок
// 	// заменить sharedApi.BlocksScoreAt на соответствующую функцию из вашего кода
// 	if err := sharedApi.BlockScoreAt(w, req); err != nil {
// 		fmt.Printf("Error calling BlocksScoreAt: %v\n", err)
// 		return 1
// 	}

// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		fmt.Printf("Expected status 200, got %d, body: %s\n", w.Code, w.Body)
// 		return 1
// 	}

// 	// Возврат кода статуса ответа HTTP
// 	return 0
// }

// func FuzzBlocksID(data []byte) int {
// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Преобразование данных JSON в структуру TestData
// 	var testData TestData
// 	if err := json.Unmarshal(data, &testData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	// Создание нового запроса GET
// 	req := httptest.NewRequest("GET", fmt.Sprintf("/go/blocks/id/%d", testData.ID), nil)

// 	// Вызов функции API и обработка ошибок
// 	// заменить sharedApi.BlocksID на соответствующую функцию из вашего кода
// 	if err := sharedApi.BlockIDAt(w, req); err != nil {
// 		fmt.Printf("Error calling BlocksID: %v\n", err)
// 		return 1
// 	}

// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		fmt.Printf("Expected status 200, got %d, body: %s\n", w.Code, w.Body)
// 		return 1
// 	}

// 	// Возврат кода статуса ответа HTTP
// 	return 0
// }

// func FuzzDebugSnapshotStateHash(data []byte) int {
// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Преобразование данных JSON в структуру TestData
// 	var testData TestData
// 	if err := json.Unmarshal(data, &testData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	// Создание нового запроса GET
// 	req := httptest.NewRequest("GET", fmt.Sprintf("/go/debug/snapshotStateHash/%d", testData.ID), nil)

// 	// Вызов функции API и обработка ошибок
// 	if err := sharedApi.snapshotStateHash(w, req); err != nil {
// 		fmt.Printf("Error calling DebugSnapshotStateHash: %v\n", err)
// 		return 1
// 	}

// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		fmt.Printf("Expected status 200, got %d, body: %s\n", w.Code, w.Body)
// 		return 1
// 	}

// 	// Возврат кода статуса ответа HTTP
// 	return 0
// }

// func FuzzBlocksByID(data []byte) int {
// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Преобразование данных JSON в структуру TestData
// 	var testData TestData
// 	if err := json.Unmarshal(data, &testData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	// Создание нового запроса GET
// 	req := httptest.NewRequest("GET", fmt.Sprintf("/go/blocks/id/%d", testData.ID), nil)

// 	// Вызов функции API и обработка ошибок
// 	if err := sharedApi.BlockIDAt(w, req); err != nil {
// 		fmt.Printf("Error calling BlocksByID: %v\n", err)
// 		return 1
// 	}

// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		fmt.Printf("Expected status 200, got %d, body: %s\n", w.Code, w.Body)
// 		return 1
// 	}

// 	// Возврат кода статуса ответа HTTP
// 	return 0
// }
