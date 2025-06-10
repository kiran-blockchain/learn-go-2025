
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/kiran-blockchain/go_chat_with_redis/handlers"
)

func main() {
    http.HandleFunc("/register", handlers.RegisterHandler)
    http.HandleFunc("/login", handlers.LoginHandler)

    fmt.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
