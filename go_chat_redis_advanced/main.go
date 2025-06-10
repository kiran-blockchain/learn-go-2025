package main

import (
	"fmt"
	"log"
	"net/http"

	"go_chat_redis_advanced/handlers"
	redisclient "go_chat_redis_advanced/redis"
)

func main() {
    http.HandleFunc("/register", handlers.RegisterHandler)
    http.HandleFunc("/login", handlers.LoginHandler)
    http.HandleFunc("/ws", handlers.WebSocketHandler)
    http.HandleFunc("/users", handlers.UsersHandler)
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    fmt.Println("Server running on http://localhost:8080")
    redisclient.Init()
    log.Fatal(http.ListenAndServe(":8080", nil))
}
