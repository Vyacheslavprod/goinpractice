// Передача ошибок по протоколу HTTP
package main

import "net/http"

func displayErr(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "An Error Occurred", http.StatusForbidden) // Возврат кода HTTP-состояния 403 с сообщением 
}

func main() {
	http.HandleFunc("/", displayErr)
	http.ListenAndServe(":8080", nil)
}