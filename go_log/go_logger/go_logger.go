// Журналирование в файл
package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./log.txt") // Создание файла журнала
	defer logfile.Close() // Гарантировать закрытие

	logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile) // Создание регистратора

	// Отсылка сообщений
	logger.Println("This is a regular message.")
	logger.Fatalln("This is a fatal error.")
}