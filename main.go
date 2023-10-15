package main

import (
	"animal-orphanage/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.InitializeAnimalsRoutes(r)

	r.Run(":8080")
}
