package main

import (
	"database/sql"
	"rest-api/config"
	"rest-api/dependency"
	"rest-api/routers"
)

func main() {
    var db *sql.DB
    defer db.Close()
    db=config.ConnectDB()

    // Build DI container
    container := dependency.BuildContainer(db)

    // Setup router
    r := routers.SetupRouter(container.UserController,container.OrderController)

    
    // Start server
    r.Run(":4000")
}
