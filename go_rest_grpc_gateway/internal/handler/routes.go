package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
    r.GET("/users/:id", GetUserHandler)
}