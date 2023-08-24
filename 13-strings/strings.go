package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// example_1()
	example_2()
}

func example_1() {
	myFile := filepath.FromSlash(filepath.Join("..", "12-files", "my_file.txt"))
	content, e := os.ReadFile(myFile)

	if e != nil {
		println("File not found")
		log.Panic(e.Error())
	}

	// str := strings.Split(string(content), " ")
	str := strings.ReplaceAll(string(content), "\n", "-")

	fmt.Println(str)
}

func example_2() {

	var result string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Type something: ")
	result, _ = reader.ReadString('\n')

	arr := strings.Split(result, " ")

	fmt.Printf("%d words in array", len(arr))

	for i, v := range arr {
		fmt.Printf("\n%d - %s", i, v)
	}

}
