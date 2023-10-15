package routes

import (
	"animal-orphanage/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeAnimalsRoutes(router *gin.Engine) {
	animalsGroup := router.Group("/animals")
	{
		animalsGroup.GET("/", controllers.GetAnimals)
		animalsGroup.GET("/:class", controllers.GetAnimalClass)
		animalsGroup.GET("/:name", controllers.GetAnimalName)
		animalsGroup.POST("/", controllers.CreateAnimal)
		animalsGroup.DELETE("/:id", controllers.DeleteAnimal)
	}
}
