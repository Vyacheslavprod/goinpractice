// Загрузка статических файлов в память и их обслуживание
package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Структура данных для хранения файла в памяти
type cacheFile struct {
	content io.ReadSeeker
	modTime time.Time
}

var cache map[string]*cacheFile // Карта файлов, хранящихся в памяти
var mutex = new(sync.RWMutex) // Мьютекс для устранения состояния гонки при параллельном внесении изменений в кэш

func main() {
	cache = make(map[string]*cacheFile) // Создание карты файлов
	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(res http.ResponseWriter, req *http.Request) {
	mutex.RLock()
	v, found := cache[req.URL.Path] // Загрузка файла из кэша, если он уже там
	mutex.RUnlock()
	if !found { // Если файл отсутсвует в кэше, загрузить его
		mutex.Lock()
		defer mutex.Unlock() // Нельзя записывать в карту сразу несколько файлов или читать их во время записи

		fileName := "./files" + req.URL.Path // Открыть файл для кэширования и отложить его закрытие
		f, err := os.Open(fileName)
		if err != nil {
			http.NotFound(res, req) // Обработка ошибок открытия файла
			return
		}
		defer f.Close()
		var b bytes.Buffer
		_, err = io.Copy(&b, f) // Копирование файла в буфер памяти
		if err != nil {
			http.NotFound(res, req) // Обработка ошибки копирования в память
			return
		}
		r := bytes.NewReader(b.Bytes()) // Поместить байты в Reader для дальнейшего использования

		// Заполнить объект и сохранить в кэше для дальнейшего использования
		info, err := f.Stat()
		if err != nil {
			http.NotFound(res, req) // Обработка ошибки получения информации о файле
			return
		}
		v := &cacheFile{
			content: r,
			modTime: info.ModTime(),
		}
		cache[req.URL.Path] = v
	}
	http.ServeContent(res, req, req.URL.Path, v.modTime, v.content) // Вернуть файл из кэша
}