// Клиент сетевого регистратора
package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1902") // Подключение к серверу журналирования
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close() // Гарантировать закрытие соединения даже в случае аварии
	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f) // Отправка регистрационных сообщений в сетевое соединение
	
	// Вывести сообщение и инициировать аварию, но не использовать Fatalln
	logger.Println("This is a regular message")
	log.Panicln("This is a panic")
}