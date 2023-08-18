package main

import "fmt"

const WORD string = "Const in global scope"

var word string

func main() {

	word = "Main scope"

	fmt.Println(word)
	fmt.Println(WORD)

	print()
}

func print() {
	word = "Print scope"

	fmt.Println(word)
	fmt.Println(WORD)
}
