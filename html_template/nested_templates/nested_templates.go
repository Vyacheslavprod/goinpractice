// Использование вложенных шаблонов
package main

import (
	"html/template"
	"net/http"
)

var t *template.Template

// Загрузка двух шаблонов в объект шаблона
func init() {
	t = template.Must(template.ParseFiles("html_template/nested_templates/index.html", "html_template/nested_templates/head.html"))
}

type Page struct {
	Title, Content string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "An Example",
		Content: "Have fun stormin da castle",
	}
	t.ExecuteTemplate(w, "index.html", p) // Обработка шаблона с передачей данных
}

func main() {
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8090", nil)
}