// Простая блокировка посредством каналов
package main

import (
	"fmt"
	"time"
)

func main() {
	lock := make(chan bool, 1)
	for i := 1; i < 7; i++ { // Вызов шести сопрограмм, совместно использующих блокирующий канал
		go worker(i, lock)
	}
	time.Sleep(10 * time.Second)
}

func worker(id int, lock chan bool) {
	fmt.Printf("%d wants the lock\n", id)
	// Рабочий процесс устанавливает блокировку, отправляя сообщение
	// Первый рабочий процесс захватывает единичный объем, что делает его собственником блокировки
	// Остальные окажутся заблокированными
	lock <- true
	fmt.Printf("%d has the lock\n", id)
	time.Sleep(500 * time.Millisecond) // Фрагмент между lock <- true и <- lock выполняется под защитой блокировки
	fmt.Printf("%d is releasing the lock\n", id)
	// Снять блокировку, прочитав значение из канала
	// В результате в буфере освободится место, и следующая функция сможет установить блокировку
	<-lock
}
