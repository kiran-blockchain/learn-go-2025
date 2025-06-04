package dependency

import (
	"github.com/kiran-blockchain/rest-api-gorm/controllers"
	"github.com/kiran-blockchain/rest-api-gorm/repositories"
	services "github.com/kiran-blockchain/rest-api-gorm/service"
	"gorm.io/gorm"
)

type Container struct {
     UserController *controllers.UserController
}

func BuildContainer(db *gorm.DB) *Container {
	userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userController  := controllers.NewUserController(userService)
	
	return &Container{
        UserController: userController,
    }
}
