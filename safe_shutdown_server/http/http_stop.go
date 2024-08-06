package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "time"
)

func main() {
    handler := newHandler()

    srv := &http.Server{
        Addr:    ":8080",
        Handler: handler,
    }

    // Create a channel to listen for OS signals
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, os.Interrupt)

    // Start the server in a goroutine
    go func() {
        fmt.Println("Starting server on :8080")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Printf("Error starting server: %v\n", err)
            os.Exit(1)
        }
    }()

    // Wait for a signal to gracefully shut down the server
    <-ch
    fmt.Println("Shutting down server...")

    // Create a context with a timeout for the graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Shut down the server gracefully
    if err := srv.Shutdown(ctx); err != nil {
        fmt.Printf("Error shutting down server: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Server stopped")
}

func newHandler() http.Handler {
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