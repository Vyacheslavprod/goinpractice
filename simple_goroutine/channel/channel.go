package main

import (
	"fmt"
	"time"
)

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <- c // Ожидание целового значения
		fmt.Print(num, " ")
	}
}

func main() {
	c := make(chan int)
	a := []int{8, 6, 7, 5, 3, 0, 9, -1}
	go printCount(c)
	for _, v := range a {
		c <- v  // Запись целового значения в канал
	}
	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")

}