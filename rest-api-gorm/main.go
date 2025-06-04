package main

import (
	"github.com/kiran-blockchain/rest-api-gorm/config"
	"github.com/kiran-blockchain/rest-api-gorm/dependency"
	"github.com/kiran-blockchain/rest-api-gorm/routers"
)

func main() {
	//var db *gorm.DB
	db := config.ConnectDB()
	//fmt.Println(db.Connection())
	//r := *gin.Default();
	// Build DI container
	container := dependency.BuildContainer(db)

	// Setup router
	r := routers.SetupRouter(container.UserController)

	// Start server
	r.Run(":4000")
}
