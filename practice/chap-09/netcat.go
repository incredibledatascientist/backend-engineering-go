package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// nc [options] host port

	// nc -l 127.0.0.1 8080 // Listener server
	// make connection to server
	// nc 127.0.0.1 8080 // Buf first start the server then only it works

	// netcat -l localhost 8080 // Server
	// netcat localhost 8080 // Client
	// nc -l // Server
	// nc -v // Verbose
	// nc -u // UDP

	connect := "localhost:8888"
	conn, err := net.Dial("tcp", connect)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(conn, "First message from tcp client.")
	
}
