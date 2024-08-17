// Обслуживание файлов
package main

import "net/http"

func main() {
	dir := http.Dir("./files")
	http.ListenAndServe(":3003", http.FileServer(dir))
}