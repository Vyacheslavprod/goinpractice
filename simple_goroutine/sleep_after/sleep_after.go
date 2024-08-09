// Реализация паузы с помощью методов Sleep и After
package main

import "time"

func main() {
	time.Sleep(5 * time.Second)
	sleep := time.After(5 * time.Second) // Создание канала для получения уведомления через 5 секунд с последующей приостановкой до поступления уведомления
	<-sleep
}