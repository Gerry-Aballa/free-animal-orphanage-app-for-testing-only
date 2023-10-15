package models

import "gorm.io/gorm"

type Sex string
type AnimalClass string

const (
	Male   Sex         = "Male"
	Female Sex         = "Female"
	Dog    AnimalClass = "Dog"
	Cat    AnimalClass = "Cat"
)

type Animal struct {
	ID    int
	Name  string
	Sex   Sex
	Class AnimalClass
	Age   int
	gorm.Model
}

var Animals []Animal
