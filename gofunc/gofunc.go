package main

import (
	"fmt"
	"time"
)

func count() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go count()
	time.Sleep(time.Second * 2)
	fmt.Println("Hello world")
	time.Sleep(time.Second * 3)
}