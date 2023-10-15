package main

import (
	"animal-orphanage/controllers"
	"animal-orphanage/database"
	"animal-orphanage/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db, err := database.SetDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	controllers.SetDB(db)

	routes.InitializeAnimalsRoutes(r)

	r.Run(":8080")
}
