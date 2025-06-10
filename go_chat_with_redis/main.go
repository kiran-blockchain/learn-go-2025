
package main

import (
    "fmt"
    "log"
    "net/http"

    "go_chat/handlers"
)

func main() {
    http.HandleFunc("/register", handlers.RegisterHandler)
    http.HandleFunc("/login", handlers.LoginHandler)

    fmt.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
