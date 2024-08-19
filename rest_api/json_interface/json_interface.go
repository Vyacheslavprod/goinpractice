// Разбор данных в формате JSON с неизвестной структурой
// Преобразование данных в формате JSON в тип interface{}
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var ks = []byte(`{
"firstName": "Jean",
"lastName": "Bartik",
"age": 86,
"education": [
	{	
		"institution": "Northwest Missouri State Teachers College",
		"degree": "Bachelor of Science in Mathematics"
	},
	{
		"institution": "University of Pennsylvania",
		"degree": "Masters in English"
	}
	],
		"spouse": "William Bartik",
"children": [
	"Timithy John Bartik",
	"Jane Helen Bartik",
	"Mary Ruth Bartik"
]
}`)

func main() {
	var f interface{}
	err := json.Unmarshal(ks, &f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(f)
	printJSON(f)

	// Доступ к firstName
	m := f.(map[string]interface{})
	fmt.Println(m["firstName"])
}

// Обход JSON-данных произвольной структуры
func printJSON(v interface{}) {
	// Для каждого значения, полученного из формата JSON, выводится информация о типе и само значение.
	// Для объектов и массивов JSON рекурсивно вызывается функция printJSON, которая выводит их внутреннее содержимое
	switch vv := v.(type) { // Обработка в зависимости от типа значения
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is an array:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	default:
		fmt.Println("Unknown type")
	}
}