// Обслуживание файлов с помощью пользовательского обработчика
package main

import "net/http"

func main() {
	http.HandleFunc("/", readme) // Регистрация обработчика для всех путей
	http.ListenAndServe(":3000", nil)
}

func readme(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "file/readme.txt") // Обслуживать содержимое файла
}