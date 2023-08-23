package main

import "fmt"

func main() {

	// if
	// age := 6
	// age := 18
	age := 60

	if age >= 18 && age < 60 {
		println("Of legal age")
	} else if age >= 60 {
		println("Elderly")
	} else {
		println("Minor")
	}

	//switch
	option := "a"

	switch option {
	case "A", "a":
		println("Choice is A")
	case "B":
		println("Choice is B")
	default:
		println("No choice")
	}

	//for
	counter := 5

	for i := 0; i < counter; i++ {
		println(i)
	}

	for x := 4; x >= 0; x-- {
		println(x)
	}

	var name []string
	name = append(name, "Rodrigo")
	name = append(name, "Carvalho")
	name = append(name, "de Brito")

	for _, v := range name {
		if v == "Rodrigo" {
			println("----------")
		}

		println(v)
	}

	for i, v := range name {
		fmt.Printf("Index: %d - Value: %s\n", i, v)
	}
}
