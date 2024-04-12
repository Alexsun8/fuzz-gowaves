import yaml

api_spec_yaml_path = 'api_spec_for_fuzz/get_smth_api.yaml'
final_file_with_tests = 'api_fuzz_tests.go'


def generate_fuzz_test(route, method):
    function_name = f"Fuzz{route.replace('/', '').title().replace('-', '')}"
    function_name = function_name.replace('{', "").replace('}',"")
    method = method.lower()
    
    route_parts = route.split('/')
    id_index = -1
    # print(route_parts)
    
    param = ""
    # Находим индекс опциональной переменной в маршруте
    for i, part in enumerate(route_parts):
        if '{' in part and '}' in part:
            id_index = i
            param = part.replace('{', "").replace('}',"")
            break
        # Инициализация переменных для опциональных параметров

    if param != "":

        if method == 'get':
            return f"""
func {function_name}(data []byte) int {{
    // Преобразование данных в JSON
    var jsonData map[string]interface{{}}
    if err := json.Unmarshal(data, &jsonData); err != nil {{
        fmt.Println("Error unmarshalling JSON data:", err)
        return -1
    }}

    
    {param} := fmt.Sprint(jsonData["{param}"])

    // Формирование маршрута с учетом опционального параметра
    reqURL := replaceRouteParams("{route}", {param}, {id_index})

    // Инициализация нового запроса GET
    req, err := http.NewRequest("GET", reqURL, nil)
    if err != nil {{
        fmt.Println("Error creating request:", err)
        return -1
    }}

    // Инициализация HTTP Recorder
	w := httptest.NewRecorder()

    // Вызов функции {function_name} и обработка ошибок
	err = sharedApi.{function_name}(w, req)
	if err != nil {{
		if containsStandardError(err) {{
            return -1
		}}

		if strings.Contains(err.Error(), "timeout waiting response from internal") {{
			fmt.Printf("Error calling {function_name}: DOS ATTACK!!!!")
			return 1
		}}

		fE := fmt.Sprintf("Error calling {function_name}: %v", err)
		fmt.Printf(fE)
		return 1
	}}
	// Проверка кода статуса ответа (пример: 200 - OK)
	if w.Code != http.StatusOK {{
		sE := fmt.Sprintf("Error: {route} Expected status 200, got %d", w.Code)
		fmt.Printf(sE)

		return 1
	}}

    // Дополнительные проверки по необходимости

    return 0
}}
"""
        elif method == 'post':
            return f"""
func {function_name}(data []byte) int {{
    // Преобразование данных в JSON
    var jsonData map[string]interface{{}}
    if err := json.Unmarshal(data, &jsonData); err != nil {{
        fmt.Println("Error unmarshalling JSON data:", err)
        return -1
    }}

    {param} := fmt.Sprint(jsonData["{param}"])

    // Формирование маршрута с учетом опционального параметра
    reqURL := replaceRouteParams("{route}", {param}, {id_index})

    // Инициализация нового запроса POST
    postData, err := json.Marshal(jsonData["postData"])
    if err != nil {{
        fmt.Println("Error marshalling JSON data:", err)
        return -1
    }}

    req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(postData))
    if err != nil {{
        fmt.Println("Error creating request:", err)
        return -1
    }}

    // Инициализация HTTP Recorder
	w := httptest.NewRecorder()

    // Вызов функции {function_name} и обработка ошибок
	err = sharedApi.{function_name}(w, req)
	if err != nil {{
		if containsStandardError(err) {{
            return -1
		}}

		if strings.Contains(err.Error(), "timeout waiting response from internal") {{
			fmt.Printf("Error calling {function_name}: DOS ATTACK!!!!")
			return 1
		}}

		fE := fmt.Sprintf("Error calling {function_name}: %v", err)
		fmt.Printf(fE)
		return 1
	}}
	// Проверка кода статуса ответа (пример: 200 - OK)
	if w.Code != http.StatusOK {{
		sE := fmt.Sprintf("Error: {route} Expected status 200, got %d", w.Code)
		fmt.Printf(sE)

		return 1
	}}

    // Дополнительные проверки по необходимости

    return 0
}}
"""
        else:
            return None
    else:
        if method == 'get':
            return f"""
func {function_name}(data []byte) int {{
    // Преобразование данных в JSON
    var jsonData map[string]interface{{}}
    if err := json.Unmarshal(data, &jsonData); err != nil {{
        fmt.Println("Error unmarshalling JSON data:", err)
        return -1
    }}

    // Формирование маршрута с учетом опционального параметра
    reqURL := "{route}"

    // Инициализация нового запроса GET
    req, err := http.NewRequest("GET", reqURL, nil)
    if err != nil {{
        fmt.Println("Error creating request:", err)
        return -1
    }}

    // Инициализация HTTP Recorder
	w := httptest.NewRecorder()

    // Вызов функции {function_name} и обработка ошибок
	err = sharedApi.{function_name}(w, req)
	if err != nil {{
		if containsStandardError(err) {{
            return -1
		}}

		if strings.Contains(err.Error(), "timeout waiting response from internal") {{
			fmt.Printf("Error calling {function_name}: DOS ATTACK!!!!")
			return 1
		}}

		fE := fmt.Sprintf("Error calling {function_name}: %v", err)
		fmt.Printf(fE)
		return 1
	}}
	// Проверка кода статуса ответа (пример: 200 - OK)
	if w.Code != http.StatusOK {{
		sE := fmt.Sprintf("Error: {route} Expected status 200, got %d", w.Code)
		fmt.Printf(sE)

		return 1
	}}

    return 0
}}
"""
        elif method == 'post':
            return f"""
func {function_name}(data []byte) int {{
    // Преобразование данных в JSON
    var jsonData map[string]interface{{}}
    if err := json.Unmarshal(data, &jsonData); err != nil {{
        fmt.Println("Error unmarshalling JSON data:", err)
        return -1
    }}

    // Формирование маршрута с учетом опционального параметра
    reqURL := "{route}"

    req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(data))
    if err != nil {{
        fmt.Println("Error creating request:", err)
        return -1
    }}

    	// Инициализация HTTP Recorder
	w := httptest.NewRecorder()

    // Вызов функции {function_name} и обработка ошибок
	err = sharedApi.{function_name}(w, req)
	if err != nil {{
		if containsStandardError(err) {{
            return -1
		}}

		if strings.Contains(err.Error(), "timeout waiting response from internal") {{
			fmt.Printf("Error calling {function_name}: DOS ATTACK!!!!")
			return 1
		}}

		fE := fmt.Sprintf("Error calling {function_name}: %v", err)
		fmt.Printf(fE)
		return 1
	}}
	// Проверка кода статуса ответа (пример: 200 - OK)
	if w.Code != http.StatusOK {{
		sE := fmt.Sprintf("Error: {route} Expected status 200, got %d", w.Code)
		fmt.Printf(sE)

		return 1
	}}

    // Дополнительные проверки по необходимости

    return 0
}}
"""


def api_tests_generator(api_spec):
    for route, methods in api_spec['paths'].items():
        for method, _ in methods.items():
            # Получаем параметры маршрута (если они есть)
            yield generate_fuzz_test(route, method)


if __name__ == '__main__':
    # Загрузка спецификации API из файла
    with open(api_spec_yaml_path, 'r') as file:
        api_spec = yaml.safe_load(file)

    # Запись тестов в файл
    with open(final_file_with_tests, 'w') as file:
        file.write('package fuzz_WE\n\n')
        file.write('import (\n')
        file.write('\t"bytes"\n')
        file.write('\t"fmt"\n')
        file.write('\t"net/http"\n')
        file.write('\t"strings"\n')
        file.write('\t"encoding/json"\n')
        file.write(')\n\n')
        file.write(f"""
func replaceRouteParams(route string, value string, param_index int) string {{
	// Создаем регулярное выражение для поиска параметров в crj,rf[]
    parts := strings.Split(route, "/")
    parts[param_index] = value
	// Заменяем все параметры на значение value
	newRoute := strings.Join(parts, "/")
	return newRoute
}}               
                   """)
        # file.write('var routeParts = strings.Split(reqURL, "/")\n')
        # file.write('var idIndex = -1\n')
        for test in api_tests_generator(api_spec=api_spec):
            file.write(f'\n{test}\n')
