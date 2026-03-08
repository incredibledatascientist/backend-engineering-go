package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr := "localhost:8080"

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("Connected to server:", addr)

	// Reader for server responses
	serverReader := bufio.NewReader(conn)

	// Reader for user input
	inputReader := bufio.NewReader(os.Stdin)

	for {
		// Read input from terminal
		fmt.Print("Enter message: ")
		text, err := inputReader.ReadString('\n')
		if err != nil {
			log.Printf("Input error: %v", err)
			continue
		}

		// Send to server
		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Printf("Write error: %v", err)
			return
		}

		// Read server response
		_, err = serverReader.ReadString('\n')
		if err != nil {
			log.Printf("Read error: %v", err)
			return
		}

		// fmt.Print("Server:", response)
	}
}
