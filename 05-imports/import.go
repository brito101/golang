package main

import (
	"fmt"
	// "log"
	// msg "log"
	. "log"
)

// go install github.com/sirupsen/logrus@latest

func main() {
	fmt.Println("Init application")

	// log.Println("Default message")
	// msg.Println("Default message")
	Println("Print from log package")

	// log.Fatal("Application Error!")
	// msg.Fatal("Application Error!")
	Fatal("Application Error!")

	fmt.Println("Continue run application")
}
