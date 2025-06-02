package controllers

import (
	"net/http"
	"rest-api/models"
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

func (c *UserController) CreateUser(ctx *gin.Context) {
   //convert input data to json 
   var user models.User
    err:= ctx.ShouldBindJSON(&user)
    id,err := c.service.CreateUser(user)

    if err!=nil{
        ctx.JSON(http.StatusInternalServerError,nil) 
    }else{
        ctx.JSON(http.StatusOK,id)
    }   
}


