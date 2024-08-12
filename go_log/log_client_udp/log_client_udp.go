// Журналирование по протоколу UDP
package main

import (
	"log"
	"net"
	"time"
)

func main() {
	timeout := 30 * time.Second // Явное добавление задержки
	conn, err := net.DialTimeout("udp", "localhost:1902", timeout) // Подключение к серверу журналирования
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