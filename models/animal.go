package models

type Sex string
type AnimalType string

const (
	Male   Sex        = "Male"
	Female Sex        = "Female"
	Dog    AnimalType = "Dog"
	Cat    AnimalType = "Cat"
)

type Animal struct {
	ID   int
	Name string
	Sex  Sex
	Type AnimalType
	Age  int
}

var Animals []Animal
