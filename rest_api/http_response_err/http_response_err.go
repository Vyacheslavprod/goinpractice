// Преобразование HTTP-ответа в ошибку
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Структура для хранения данных об ошибке
type Error struct {
	HTTPCode int `json:"-"`
	Code int `json:"code,omitempty"`
	Message string `json:"message"`
}

// Метод Error реализует интерфейс error и опирается на структуру Error
func (e Error) Error() string {
	fs := "HTTP: %d, Code: %d, Message: %s"
	return fmt.Sprintf(fs, e.HTTPCode, e.Code, e.Message)
}

// Для выполнения запросов вместо функции http.Get должна использоваться функция get
func get(u string) (*http.Response, error) {
	// Использование http.Get для плучения ресурса и возврата любых ошибок http.Get errors
	res, err := http.Get(u)
	if err != nil {
		return res, err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		// Проверка типа содержимого ответа и возврат ошибки, если он неправильный
		if res.Header.Get("Content-Type") != "application/json" {
			sm := "unknown error. HTTP status : %s"
			return res, fmt.Errorf(sm, res.Status)
		}
		// Чтение тела ответа в буфер
		b, _ := io.ReadAll(res.Body)
		res.Body.Close()
		// Разобрать JSON-ответ, записать данные в структуру и вернуть ошибку
		var data struct {
			Err Error `json:"error"`
		}
		err = json.Unmarshal(b, &data)
		if err != nil {
			sm := "unable to parse json: %s. HTTP status: %s"
			return res, fmt.Errorf(sm, err, res.Status)
		}
		// Добавление кода HTTP-состояния в экземпляр Error
		data.Err.HTTPCode = res.StatusCode

		// Возврат пользовательской ошибки с ответом
		return res, data.Err
	}
	return res, nil
}

func main() {
	res, err := get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, _ := io.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}