// Обработка выгрузки одного файла
package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("file.html")
		t.Execute(w, nil)
	} else {
		f, h, err := r.FormFile("file") // Получение указателя на файл, заголовочную инфо. и ошибку для формы поля
		if err != nil {
			panic(err)
		}
		defer f.Close()

		dir := "./files" // Измените путь на относительный
  		err = os.MkdirAll(dir, os.ModePerm) // Создание директории
  		if err != nil {
   			http.Error(w, "Unable to create directory", http.StatusInternalServerError)
   			return
  		}

		
		filename := dir + "/" + h.Filename
		out, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		io.Copy(out, f)
		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm) // Регистрация обработчика для всех путей
	http.ListenAndServe(":3000", nil)
}