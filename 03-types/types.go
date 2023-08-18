package main

import "fmt"

// Global vars don't accept dynamic inference
// name := "Rodrigo" //error
var name = "Rodrigo"

func main() {
	var a string = "String"
	var b int = 2
	// var c float32 = 2.5
	var c float64 = 2.5
	var d bool = true
	var e [3]string
	e[0] = "Rodrigo"
	e[1] = "Carvalho"
	e[2] = "de Brito"

	f := "Dynamic" //Dynamic inference using ":"

	fmt.Println("String: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float: ", c)
	fmt.Println("Boolean: ", d)
	fmt.Println("Array: ", e)
	fmt.Println("Dynamic inference: ", f)

	print()

}

func print() {
	fmt.Println(name)
}
