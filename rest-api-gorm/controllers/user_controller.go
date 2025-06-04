package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/kiran-blockchain/rest-api-gorm/service"
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
