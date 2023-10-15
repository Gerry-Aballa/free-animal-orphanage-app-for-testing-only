package controllers

import (
	"animal-orphanage/database"
	"animal-orphanage/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

func GetAnimals(c *gin.Context) {
	//TODO: implement this method to return all animals in the database
	var animals []models.Animal
	result := database.DB.Find(&animals)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch contacts"})
		return
	}

	c.JSON(200, animals)
}

func GetAnimalClass(c *gin.Context) {
	//TODO: implement this method to get animal by type from the database
}

func GetAnimalName(c *gin.Context) {
	//TODO: implement this method to get an animal by name from the database
}

func CreateAnimal(c *gin.Context) {
	//TODO: implement this method to create a new animal in the database
	var animal models.Animal
	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := database.DB.Create(&animal)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create animal"})
		return
	}

	c.JSON(201, animal)
}

func DeleteAnimal(c *gin.Context) {
	//TODO: implement this method to delete an animal from the database
}
