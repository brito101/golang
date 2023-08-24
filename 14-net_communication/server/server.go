package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listener, e := net.Listen("tcp", "0.0.0.0:9090")

	if e != nil {
		log.Fatal("Error in config listener: ", e.Error())
	}

	con, err := listener.Accept()

	if err != nil {
		log.Fatal("Error in start listener: ", err.Error())
	}

	for {
		message, _ := bufio.NewReader(con).ReadString('\n')
		con.Write([]byte(message))
	}

}
