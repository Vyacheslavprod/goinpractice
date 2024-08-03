package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://example.com") // Создание HTTP - GET запроса
	body, _ := io.ReadAll(resp.Body) // Чтение тела ответа
	fmt.Println(string(body)) // Вывод тела в виде строки
	resp.Body.Close() // Закрытие соединения
}