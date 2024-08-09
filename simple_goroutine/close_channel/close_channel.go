// Использование завершающего канала
package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	done := make(chan bool) // Дополнительный канал с типом данных bool для сообщения о завершении
	until := time.After(5 * time.Second)
	go send(msg, done)
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			done <- true // По истечении заданного интервала времени сообщить сопрограмме send, что работа завершена
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

func send(ch chan <-string, done <-chan bool) { // ch используется для отправки, а done для получения
	for {
		select {
		case <-done:
			println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}