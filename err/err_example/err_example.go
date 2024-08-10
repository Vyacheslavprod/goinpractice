// Возврат ошибки
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Функция Concat объединяет строки, разделяя их пробелами.
// Она возвращает пустую строку и ошибку, если не получила ни одной строки.
func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("no strings supplied") // Вернуть ошибку, если ничего не было передано
	}
	return strings.Join(parts, " "), nil // Вернуть новую строку и значение nil
}

func main() {
	args := os.Args[1:]
	result, _ := Concat(args...)
	fmt.Printf("concatenated string: '%s'\n", result)
	/*
	При возвращение полезного значения, можно опустить обработку ошибки
	if result, err := Concat(args...); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("concatenated string: '%s'\n", result)
	}
	*/
}