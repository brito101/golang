package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	con, e := net.Dial("tcp", "127.0.0.1:9090")

	if e != nil {
		log.Fatal("Error in connect server: ", e.Error())
	}

	for {
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		con.Write([]byte(message))
		res, _ := bufio.NewReader(con).ReadString('\n')
		fmt.Println("Server response: ", res)
	}
}
