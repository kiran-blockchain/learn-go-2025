package routers

import (
	"rest-api/controllers"
	middlewares "rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userController *controllers.UserController,orderController *controllers.OrderController) *gin.Engine {
    r := gin.Default()

    // Users endpoint
    r.POST("/register", userController.CreateUser)
    r.POST("/login", userController.Login,userController.GetAllUsers)

         // Protected endpoints
    auth := r.Group("/auth")
    //intercept the JWT token and if it valid grant access to the below routes
    auth.Use(middlewares.JWTAuthMiddleware())
    auth.GET("/getusers", userController.GetAllUsers)


    r.POST("/orders", orderController.PlaceOrder)

    return r
}
