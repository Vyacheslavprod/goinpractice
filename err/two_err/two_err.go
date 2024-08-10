// Обработка двух различных ошибок
package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var ErrTimeout = errors.New("the request timed out") // Создание экземпляра ошибки превышения времени ожидания
var ErrRejected = errors.New("the request was rejected") // Создание экземпляра ошибки отказа
var random = rand.New(rand.NewSource(35))

func main() {
	response, err := SendRequest("Hello")
	// Обработка превышения времени ожидания повторением попытки
	for err == ErrTimeout {
		fmt.Println("Timeout. Retrying")
		response, err = SendRequest("Hello")
	}
	// Обработка любой другой ошибки как сбоя
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response) // Если нет ошибок, напечатать результат
	}
}

func SendRequest(req string) (string, error) {
	switch random.Int() % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", ErrRejected
	default:
		return "", ErrTimeout
	}
}