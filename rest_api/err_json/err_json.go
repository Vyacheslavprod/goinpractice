// Ответ с сообщением об ошибке в формате JSON
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Тип для хранения информации об ощибке, включая метаданные о структуре JSON
type Error struct {
	HTTPCode int `json:"-"`
	Code int `json:"code,omitempty"`
	Message string `json:"massage"`
}

// Функция JSONError похожа на функцию http.Error, но тело ответа имеет формат JSON
func JSONError(w http.ResponseWriter, e Error) {
	data := struct { // Обертывание структуры Error анонимной структурой со свойством error
		Err Error `json:"error"`
	}{e}
	
	// Обертывание структуры Error анонимной структурой со свойством error
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// Установка MIME-типа ответа в application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode) // Гарантирование правильной установки кода ошибки HTTP
	fmt.Fprint(w, string(b)) //Вывод тела в формате JSON
}

func diplayError(w http.ResponseWriter, r *http.Request) {
	// Создание экземпляра Error для использования в ответе с описанием ошибки
	e := Error{
		HTTPCode: http.StatusForbidden,
		Code: 123,
		Message: "An Error Occurred",
	}
	JSONError(w, e) // Возврат сообщения об ошибке в формате JSON при вызове обработчика HTTP
}

func main() {
	http.HandleFunc("/", diplayError)
	http.ListenAndServe(":8080", nil)
}