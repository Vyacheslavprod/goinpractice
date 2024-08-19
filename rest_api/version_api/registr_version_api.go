// Версии прикладных программных интерфейсов в URL-адресе
// Регистрация пути к API с включением версии
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type testMessage struct {
	Message string `json:"message"`
}

// Пример функции, возвращающей ответ в формате JSON
func diplayTest(w http.ResponseWriter, r *http.Request) {
	data := testMessage{"A test message."}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}

func main() {
	http.HandleFunc("/api/v1/test", diplayTest) // Включение версии API в путь, который отображается в функции-обработчик
	http.ListenAndServe(":8080", nil)
}