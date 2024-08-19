// Обращение к API v2
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	ct := "application/vnd.mytodos.json; version=2.0" // Тип содержимого с номером версии API
	req, _ := http.NewRequest("GET", "http://localhost:8080/api", nil)
	req.Header.Set("Accept", ct) // Добавление нужного типа содержимого в запрос
	res, _ := http.DefaultClient.Do(req) // Выполнение запроса

	// Проверка типа содержимого в ответе
	if res.Header.Get("Content-Type") != ct {
		fmt.Println("Unexpected content type returned")
		return
	}

	b, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Printf("%s", b)
}