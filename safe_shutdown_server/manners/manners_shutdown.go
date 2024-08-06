package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"github.com/braintree/manners"
)

func main() {
	handler := newHandler()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go listenForShutdown(ch)

	manners.ListenAndServe(":8080", handler)
}

func newHandler() *handler {
	return &handler{}
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name != "" {
		name = "Inigo Montoya"
	}
	fmt.Fprint(w, "Hello, my name is ", name)
}

func listenForShutdown(ch <- chan os.Signal) {
	<-ch
	manners.Close()
}