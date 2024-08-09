// Нескоько функций-обработчиков
package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye", goodbye)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request)  {
	// Получение имени из строки запроса
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(w, "Hello, my name is ", name)
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	// Выборка имени из строки запроса
	path := r.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Inigo Motoya"
	}
	fmt.Fprint(w, "Goodbye ", name)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// Проверка соответствия пути домашней или неопознанной странице
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	fmt.Fprint(w, "The homepage")
}