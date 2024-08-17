// Кэширование разобранного шаблона
package main

import (
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseFiles("html_template/simple_template/simple.html")) // Синтаксический разбор при инициализации пакета

type Page struct {
	Tittle, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Tittle: "An Example",
		Content: "Have fun stormin da castle",
	}
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}