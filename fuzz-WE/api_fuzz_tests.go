package fuzz_WE

// func replaceRouteParams(route string, value string, param_index int) string {
// 	// Создаем регулярное выражение для поиска параметров в crj,rf[]
// 	parts := strings.Split(route, "/")
// 	parts[param_index] = value
// 	// Заменяем все параметры на значение value
// 	newRoute := strings.Join(parts, "/")
// 	return newRoute
// }

// func FuzzGoblocksscoreatId(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	id := fmt.Sprint(jsonData["id"])

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := replaceRouteParams("/go/blocks/score/at/{id}", id, 5)

// 	// Инициализация нового запроса GET
// 	req, err := http.NewRequest("GET", reqURL, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return -1
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzGoblocksscoreatId и обработка ошибок
// 	err = sharedApi.BlockScoreAt(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzGoblocksscoreatId: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzGoblocksscoreatId: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /go/blocks/score/at/{id} Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }

// func FuzzGoblocksidId(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	id := fmt.Sprint(jsonData["id"])

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := replaceRouteParams("/go/blocks/id/{id}", id, 4)

// 	// Инициализация нового запроса GET
// 	req, err := http.NewRequest("GET", reqURL, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return -1
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzGoblocksidId и обработка ошибок
// 	err = sharedApi.BlockIDAt(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzGoblocksidId: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzGoblocksidId: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /go/blocks/id/{id} Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }

// func FuzzGowalletaccounts(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := "/go/wallet/accounts"

// 	// Инициализация нового запроса GET
// 	req, err := http.NewRequest("GET", reqURL, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return -1
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzGowalletaccounts и обработка ошибок
// 	err = sharedApi.WalletAccounts(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzGowalletaccounts: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzGowalletaccounts: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /go/wallet/accounts Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	return 0
// }

// func FuzzBlocksId(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	id := fmt.Sprint(jsonData["id"])

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := replaceRouteParams("/blocks/{id}", id, 2)

// 	// Инициализация нового запроса GET
// 	req, err := http.NewRequest("GET", reqURL, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return -1
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzBlocksId и обработка ошибок
// 	err = sharedApi.BlockHeadersID(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzBlocksId: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzBlocksId: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /blocks/{id} Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }

// func FuzzBlocksheadersatHeight(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	height := fmt.Sprint(jsonData["height"])

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := replaceRouteParams("/blocks/headers/at/{height}", height, 4)

// 	// Инициализация нового запроса GET
// 	req, err := http.NewRequest("GET", reqURL, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return -1
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzBlocksheadersatHeight и обработка ошибок
// 	err = sharedApi.BlocksHeadersAt(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzBlocksheadersatHeight: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzBlocksheadersatHeight: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /blocks/headers/at/{height} Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }

// func FuzzBlocksheadersId(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	id := fmt.Sprint(jsonData["id"])

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := replaceRouteParams("/blocks/headers/{id}", id, 3)

// 	// Инициализация нового запроса GET
// 	req, err := http.NewRequest("GET", reqURL, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return -1
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzBlocksheadersId и обработка ошибок
// 	err = sharedApi.BlockHeadersID(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzBlocksheadersId: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzBlocksheadersId: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /blocks/headers/{id} Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }

// func FuzzAssetsdetailsId(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	id := fmt.Sprint(jsonData["id"])

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := replaceRouteParams("/assets/details/{id}", id, 3)

// 	// Инициализация нового запроса GET
// 	req, err := http.NewRequest("GET", reqURL, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return -1
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzAssetsdetailsId и обработка ошибок
// 	err = sharedApi.AssetsDetailsByID(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzAssetsdetailsId: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzAssetsdetailsId: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /assets/details/{id} Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }

// func FuzzAssetsdetails(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := "/assets/details"

// 	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(data))
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return -1
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzAssetsdetails и обработка ошибок
// 	err = sharedApi.AssetsDetailsByIDsPost(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzAssetsdetails: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzAssetsdetails: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /assets/details Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }

// func FuzzAliasbyAliasAlias(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	alias := fmt.Sprint(jsonData["alias"])

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := replaceRouteParams("/alias/by-alias/{alias}", alias, 3)

// 	// Инициализация нового запроса GET
// 	req, err := http.NewRequest("GET", reqURL, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return -1
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzAliasbyAliasAlias и обработка ошибок
// 	err = sharedApi.AddrByAlias(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzAliasbyAliasAlias: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzAliasbyAliasAlias: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /alias/by-alias/{alias} Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }

// func FuzzTransactionsinfoId(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	id := fmt.Sprint(jsonData["id"])

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := replaceRouteParams("/transactions/info/{id}", id, 3)

// 	// Инициализация нового запроса GET
// 	req, err := http.NewRequest("GET", reqURL, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return -1
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzTransactionsinfoId и обработка ошибок
// 	err = sharedApi.TransactionInfo(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzTransactionsinfoId: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzTransactionsinfoId: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /transactions/info/{id} Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }

// func FuzzPeersconnect(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := "/peers/connect"

// 	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(data))
// 	if err != nil {
// 		{
// 			fmt.Println("Error creating request:", err)
// 			return -1
// 		}
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzPeersconnect и обработка ошибок
// 	err = sharedApi.PeersConnected(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzPeersconnect: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzPeersconnect: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /peers/connect Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }

// func FuzzPeersclearblacklist(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := "/peers/clearblacklist"

// 	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(data))
// 	if err != nil {
// 		{
// 			fmt.Println("Error creating request:", err)
// 			return -1
// 		}
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzPeersclearblacklist и обработка ошибок
// 	err = sharedApi.PeersBlackListed(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzPeersclearblacklist: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzPeersclearblacklist: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /peers/clearblacklist Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }

// func FuzzDebugrollbackToId(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	id := fmt.Sprint(jsonData["id"])

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := replaceRouteParams("/debug/rollback-to/{id}", id, 3)

// 	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(data))
// 	if err != nil {
// 		{
// 			fmt.Println("Error creating request:", err)
// 			return -1
// 		}
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzDebugrollbackToId и обработка ошибок
// 	err = sharedApi.RollbackTo(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzDebugrollbackToId: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzDebugrollbackToId: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /debug/rollback-to/{id} Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }

// func FuzzEthabiAddress(data []byte) int {
// 	// Преобразование данных в JSON
// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		fmt.Println("Error unmarshalling JSON data:", err)
// 		return -1
// 	}

// 	address := fmt.Sprint(jsonData["address"])

// 	// Формирование маршрута с учетом опционального параметра
// 	reqURL := replaceRouteParams("/eth/abi/{address}", address, 3)

// 	// Инициализация нового запроса GET
// 	req, err := http.NewRequest("GET", reqURL, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return -1
// 	}

// 	// Инициализация HTTP Recorder
// 	w := httptest.NewRecorder()

// 	// Вызов функции FuzzEthabiAddress и обработка ошибок
// 	err = sharedApi.EthereumDAppABI(w, req)
// 	if err != nil {
// 		if containsStandardError(err) {
// 			return -1
// 		}

// 		if strings.Contains(err.Error(), "timeout waiting response from internal") {
// 			fmt.Printf("Error calling FuzzEthabiAddress: DOS ATTACK!!!!")
// 			return 1
// 		}

// 		fE := fmt.Sprintf("Error calling FuzzEthabiAddress: %v", err)
// 		fmt.Printf(fE)
// 		return 1
// 	}
// 	// Проверка кода статуса ответа (пример: 200 - OK)
// 	if w.Code != http.StatusOK {
// 		sE := fmt.Sprintf("Error: /eth/abi/{address} Expected status 200, got %d", w.Code)
// 		fmt.Printf(sE)

// 		return 1
// 	}

// 	// Дополнительные проверки по необходимости

// 	return 0
// }
