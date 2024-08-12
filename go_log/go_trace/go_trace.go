// Вывод трассировки стека в поток стандартного вывода
package main

import (
	"fmt"
	"runtime"
	// "runtime/debug"
)

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	// debug.PrintStack()
	buf := make([]byte, 1024) // Создание буфера
	runtime.Stack(buf, false) // Запись стека в буффер
	fmt.Printf("Trace:\n %s\n", buf)
}