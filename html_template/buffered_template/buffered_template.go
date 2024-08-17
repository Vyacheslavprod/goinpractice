// Буферизация вывода шаблона
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template
func init() {
	t = template.Must(template.ParseFiles("simple.html")) // Синтаксический разбор при инициализации пакета
}

type Page struct {
	Tittle, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Tittle: "An Example",
		Content: "Have fun stormin da castle",
	}
	var b bytes.Buffer // Создание буфера для сохранения результатов обработки шаблона
	err := t.Execute(&b, p)
	if err != nil {
		fmt.Fprintf(w, "A error occured.")
		return
	}
	b.WriteTo(w) // Вывод буферизированного ответа
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}