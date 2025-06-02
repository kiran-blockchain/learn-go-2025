package controllers

import (
	"net/http"
	"rest-api/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    service services.UserService
}

func NewUserController(s services.UserService) *UserController {
    return &UserController{service: s}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
    users,err := c.service.GetAllUsers()
    if err!=nil{
        ctx.JSON(http.StatusInternalServerError,nil)
    } else {
    ctx.JSON(http.StatusOK, users)
    }
}


