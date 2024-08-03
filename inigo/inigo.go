// Вывод Hello world через веб-сервер
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hello, my name is Inigo")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)
}