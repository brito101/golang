package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {

	file := filepath.FromSlash(filepath.Join(os.TempDir(), "file.txt"))
	// os.Create(os.TempDir() + "/" + "my_file.txt")
	// os.Create("my_file.txt")
	// f, e := os.OpenFile("my_file.txt", os.O_APPEND|os.O_WRONLY, 0660)
	// f, e := os.OpenFile("my_file.txt", os.O_TRUNC|os.O_WRONLY, 0660)
	f, e := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0660)

	if e != nil {
		println("File not found")
		log.Panic(e.Error())
	}

	defer f.Close()

	f.Write([]byte("TEST\n"))
}
