// Добавление функций для шаблонов
package main

import (
	"html/template"
	"net/http"
	"time"
)

// Применение команды dateFormat к Date
var tpl = `<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Date Example</title>
    </head>
    <body>
        <p>{{.Date | dateFormat "Jan 2, 2006"}}</p>
    </body>
</html>`

// Отображение функций Go в функции для шаблона
var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

// Функция для преобразования времени в форматированную строку
func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
	t := template.New("date") // Создание нового экземпляра template.Template
	t.Funcs(funcMap) // Передача карты с дополнительными функциями механизму шаблонов
	t.Parse(tpl) // Синтаксический разбор строки шаблона
	data := struct{ Date time.Time} {
		Date: time.Now(),
	}
	t.Execute(res, data) // Отправка шаблона и набора данных с помощью веб-сервера
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}


/*
Если одну и туже функцию потребуется использовать в нескольких шаблонах, 
можно написать отдельную функцию, которая будет создавать шаблоны и добавлять в них новые функции:
func parseTemplateString(name, tpl string) *template.Template {
	t := template.New(name)
	t.Funcs(funcMap)
	t = template.Must(t.Parse(tpl))
	return t
}
*/