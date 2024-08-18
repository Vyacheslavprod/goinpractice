// Определение вревышения времени ожидания в сети по ошибкам
package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// Функция, возвращающая true, если ошибка вызвана превышением времени ожидания
func hasTimeOut(err error) bool {
	switch err := err.(type) {
	case *url.Error:
		if err, ok := err.Err.(net.Error); ok && err.Timeout() {
			return true
		}
	// Проверить, не обнаружил ли ошибку превышения времени ожидания пакет net
	case net.Error:
		if err.Timeout() {
			return true
		}
	/*
	case *net.OpError:
		if err.Timeout() {
			return true
		}
	*/
	}
	errTxt := "use of closed network connection"
	if err != nil && strings.Contains(err.Error(), errTxt) {
		return true
	}
	return false
}

func main() {
	res, err := http.Get("http://exa") // Выполнение GET запроса
	if err != nil && hasTimeOut(err) {
		fmt.Println("A timeout error occured")
		return
	}

	b, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Printf("%s", b)
}