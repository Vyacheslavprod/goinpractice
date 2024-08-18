// Использование функции http.Get
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, _ := http.Get("http://example.com") // Выполнение GET запроса
	b, _ := io.ReadAll(res.Body) // Чтение поля ответа и закрытие объекта Body после завершения чтения
	res.Body.Close()
	fmt.Printf("%s", b) // Вывод тела в поток стандартного вывода
	fmt.Println(res.Status)
	fmt.Println(res.StatusCode)
}
