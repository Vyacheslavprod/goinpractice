// Преобразование данных в формате JSON
package main

import (
	"encoding/json"
	"fmt"
)

// Структура, представляющая информацию в формате JSON.
// Тег json отображает свойство Name в поле name в данных в формате JSON
type Person struct {
	Name string `json:"name"`
}

// Строка с данными в формате JSON
var JSON = `{
	"name": "Miracle Max"
}`

func main() {
	var p Person // Экземпляр структуры Person для сохранения данных в формате JSON
	err := json.Unmarshal([]byte(JSON), &p) // Преобразование данных в формате JSON в экземпляр структуры Person
	if err != nil {
		fmt.Println(err)
		return
	}
	// Использование заполненного объекта Person, в данном случае его вывод
	fmt.Println(p)
}