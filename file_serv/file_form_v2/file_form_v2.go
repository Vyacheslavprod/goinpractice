// Выгрузка нескольких файлов с помощью формы
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
		t, _ := template.ParseFiles("file_multiple.html")
		t.Execute(w, nil)
	} else {
		err := r.ParseMultipartForm(16 << 20)
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		data := r.MultipartForm
		files := data.File["files"]
		for _, fh := range files {
			f, err := fh.Open()
			
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			defer f.Close()

			out, err := os.Create("/tmp/" + fh.Filename)
			

			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			defer out.Close()

			_, err = io.Copy(out, f)
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
		}
		fmt.Fprint(w, "Upload complete")
	}
}

func main() {
	http.HandleFunc("/", fileForm) // Регистрация обработчика для всех путей
	http.ListenAndServe(":3000", nil)
}