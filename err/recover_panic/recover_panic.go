// Восстановление после аварии
package main

import (
	"errors"
	"fmt"
)

func main() {
	// Подготовка отложенного вызова замыкания для восстановления после аварии
	defer func () {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
	}()
	yikes() // Вызов функции возбуждающей аварию
}

func yikes() {
	panic(errors.New("something bad happened"))
}