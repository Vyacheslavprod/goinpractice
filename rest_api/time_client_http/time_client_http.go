// Простой пользовательский HTTP - клиент
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	cc := &http.Client{Timeout: time.Second}
	res, err := cc.Get("http://goinpracticebook.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, _ := io.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}