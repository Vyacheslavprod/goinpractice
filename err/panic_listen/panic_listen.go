// Обработка аварий в сопрограмме
package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":1026")
	if err != nil {
		fmt.Println("Failed to open port on 1026")
		return
	}
	// Прием новых клиентских запросов и обработка ошибок подключения
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}
		go handle(conn) // При появлении запроса передать его в функцию handle
	}
}

func handle(conn net.Conn) {
	// Отложенная функция обрабатывает аварию 
	// и гарантирует закрытие соединения в любом случае
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal error: %s", err)
		}
		conn.Close()
	}()
	// Попытка чтения строки из подключения
	reader := bufio.NewReader(conn)
	data, err := reader.ReadBytes('\n')
	// В случае ошибки чтения строки вывести сообщение и закрыть подключение
	if err != nil {
		fmt.Println("Failed to read from socket.")
		conn.Close()
	}
	response(data, conn) // При получении строки передать ее в функцию response
}

// Записать данные в сокет для отправки клиенту
func response(data []byte, conn net.Conn) {
	conn.Write(data)
	panic(errors.New("pretend I'm a real error"))
}