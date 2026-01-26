package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println("TCP Listening ...")
	l, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()
	// This will handle single connection
	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		data := string(buf[:n])
		msg := strings.TrimSpace(data)
		fmt.Print(msg)

		if msg == "STOP" {
			conn.Write([]byte("Server is closing now..."))
			fmt.Println("Server got STOP signal.")
			return
		}
	}
}
