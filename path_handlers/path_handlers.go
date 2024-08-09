// Анализ URL - адресов с помощью пакета path
package main

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func main() {
	pr := newPathResolver() // Получение экземпляра маршрутизатора
	pr.Add("GET /hello", hello)
	pr.Add("* /goodbye/*", goodbye)
	http.ListenAndServe(":8080", pr)
}

// Создание нового инициализированного объекта pathResolver
func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

// Добавление путей для внутреннего поиска
func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}

func (p *pathResolver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	check := r.Method + " " + r.URL.Path // Объединение метода и пути для проверки

	for pattern, handlerFunc := range p.handlers { // Обход зарегистрированных путей
		if ok, err := path.Match(pattern, check); ok && err == nil { // Проверка соответствия текущего пути одному из зарегистрированных
			handlerFunc(w, r) // Вызов функции для обработки пути
			return
		} else if err != nil {
			fmt.Fprint(w, err)
		}
	}
	http.NotFound(w,r)
}

func hello(w http.ResponseWriter, r *http.Request)  {
	// Получение имени из строки запроса
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(w, "Hello, my name is ", name)
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	// Выборка имени из строки запроса
	path := r.URL.Path
	parts := strings.Split(path, "/")
	name := parts[2]
	if name == "" {
		name = "Inigo Motoya"
	}
	fmt.Fprint(w, "Goodbye ", name)
}