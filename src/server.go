package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println("loading server")

	//listen
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	//accept

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	//run loop forever

	for {
		//listen message to process in newline
		message, _ := bufio.NewReader(conn).ReadString('\n')
		//output message receired
		fmt.Print("message receired:", string(message))
		//sample process received
		newMessage := strings.ToUpper(message)
		//send message client
		conn.Write([]byte(newMessage + "\n"))
	}

}
