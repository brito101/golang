package main

import (
	"fmt"
	"structures/models"
)

var person = models.Person{}

func main() {
	person.Name = "Rodrigo"
	person.Address = "Brazil"
	person.Age = 38
	person.Smoker = false

	personDetail()
}

func personDetail() {
	fmt.Println("Name: ", person.Name)
	fmt.Println("Address:", person.Address)
	fmt.Println("Age: ", person.Age)
	fmt.Println("Smoker: ", person.Smoker)
}
