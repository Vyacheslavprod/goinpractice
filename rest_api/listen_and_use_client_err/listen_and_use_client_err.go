// Чтение и использование пользовательских сообщений об ошибках
package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, _ := http.Get("http://example.com")
	switch {
	case 300 <= res.StatusCode && res.StatusCode < 400:
		fmt.Println("Redirect message")
	case 400 <= res.StatusCode && res.StatusCode < 500:
		fmt.Println("Client error")
	case 500 <= res.StatusCode && res.StatusCode < 600:
		fmt.Println("Server error")
	}
}