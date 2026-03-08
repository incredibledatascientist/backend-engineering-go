package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	addr := "localhost:8080"
	fmt.Println("server is listening on ", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Read error: %v", err)
			return
		}

		fmt.Print("Client <-", message)

		ackMsg := strings.ToUpper(strings.TrimSpace(message))
		response := fmt.Sprintf("ACK: %s\n", ackMsg)
		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Printf("Server write error: %v", err)
		}
	}
}
