package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":192.168.51.100:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		//read message to server
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("text to send:")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		//send to socket
		fmt.Fprintf(conn, text+"\n")
		//listen reply
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("message from server", message)
	}
}
