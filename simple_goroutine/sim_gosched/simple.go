package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("1. Outside a goroutine.")
	go func() {
		fmt.Println("2.Go() Inside a goroutine")
	}()
	fmt.Println("3. Outside again")

	runtime.Gosched()
}