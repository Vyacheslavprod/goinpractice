// Передача версии API в типе содержимого
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api", diTest) // Регистрация пути с несколькими типами содержимого
	http.ListenAndServe(":8080", nil)
}

func diTest(w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get("Accept") // Определение зарегистрированного ранее типа содержимого
	var err error
	var b []byte
	var ct string
	// Конструирование разных ответов, в зависимости от запрошенного типа содержимого
	switch t {
	case "application/vnd.mytodos.json; version=2.0":
		data := testMessageV2{"Version 2"}
		b, err = json.Marshal(data)
		ct = "application/vnd.mytodos.json; version=2.0"
	case "application/vnd.mytodos.json; version=1.0":
		fallthrough
	default:
		data := testMessageV1{"Version 1"}
		b, err = json.Marshal(data)
		ct = "application/vnd.mytodos.json; version=1.0"
	}
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", ct) // Установить тип содержимого в соответствии с запрошенным типом
	fmt.Fprint(w, string(b)) // Вернуть ответ клиенту
}

type testMessageV1 struct {
	Message string `json:"message"`
}

type testMessageV2 struct {
	Info string `json:"info"`
}