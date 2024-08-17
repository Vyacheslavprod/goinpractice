// Использование наследования шаблонов
package main

import (
	"html/template"
	"net/http"
)

var t map[string]*template.Template // Ассоциативный массив для хранения шаблонов с их именами
func init() {
	t = make(map[string]*template.Template) // Настройка карты шаблонов

	// Загрузка шаблонов вместе с основным шаблоном
	temp := template.Must(template.ParseFiles("templates/base.html", "templates/user.html"))
	t["user.html"] = temp
	temp = template.Must(template.ParseFiles("templates/base.html", "templates/page.html"))
	t["page.html"] = temp
}

type Page struct {
	Title, Content string
}

type User struct {
	Username, Name string
}

func displayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "Inherit",
		Content: "Have fun stormin da castle",
	}
	t["page.html"].ExecuteTemplate(w, "base", p) // Подключение шаблона к странице
}

func displayUser(w http.ResponseWriter, r *http.Request) {
	u := &User{
		Username: "swordsmith",
		Name: "Inigo Montoya",
	}
	t["user.html"].ExecuteTemplate(w, "base", u)
}

func main() {
	http.HandleFunc("/user", displayUser)
	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}