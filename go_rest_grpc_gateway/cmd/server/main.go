package main

import (
    "github.com/gin-gonic/gin"
    "go_rest_grpc_gateway/internal/handler"
)

func main() {
    r := gin.Default()
    handler.RegisterRoutes(r)
    r.Run(":8080")
}