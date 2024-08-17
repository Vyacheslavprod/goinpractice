// Использование простого HTML - шаблона
package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Tittle, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Tittle: "An Example",
		Content: "Have fun stormin da castle",
	}
	t := template.Must(template.ParseFiles("html_template/simple_template/simple.html"))
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}