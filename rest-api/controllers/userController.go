package controllers

import (
	"net/http"
	"rest-api/models"
	"rest-api/services"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	} else {
		token, isExists := ctx.Get("token")
		if isExists {
			ctx.JSON(http.StatusOK, gin.H{
				"jwt":   token,
				"users": users,
			})
		} else {
            ctx.JSON(http.StatusOK, gin.H{
				"jwt":   "",
				"users": users,
			})
		}
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	//convert input data to json
	var user models.User
	//convert the payload from postman to model
	err := ctx.ShouldBindJSON(&user)
	user.PasswordHash, err = utils.HashPassword(user.PasswordHash)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	id, err := c.service.CreateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	} else {
		ctx.JSON(http.StatusOK, id)
	}
}

func (c *UserController) Login(ctx *gin.Context) {
	//convert input data to json
	var user models.User
	//convert the payload from postman to model
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	}
	token, err := c.service.Login(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
	} else {
		ctx.Set("token", token)
		ctx.Next()
		// ctx.JSON(http.StatusOK,result)
	}
}
