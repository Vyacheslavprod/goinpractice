// Использование нескольких каналов
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := time.After(30 * time.Second) // Создание канала, посылающего сообщение по истечении 30 секунд
	echo := make(chan []byte) // Создание нового канала для передачи байтов из Stdin в Stdout

	go readStdin(echo) // Запуск сопрограммы для чтения данных из Stdin и передачи их в новый канал
	
	// Использование select для передачи данных из Stdin в Stdout, если имеются, и для завершения по событию таймера
	for {
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed out")
			os.Exit(0)
		}
	}
}

// Принимает канал для записи и отправляет в канал введенные данные
func readStdin(out chan <- []byte) {
	// Копирование данных из Stdin в объект data
	for {
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			out <- data
		}
	}
}