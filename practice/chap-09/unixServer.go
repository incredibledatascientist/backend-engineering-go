package main

import (
	"fmt"
	"net"
	"os"
)

func echo(c net.Conn) {
	for {
		buf := make([]byte, 128)
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println("Error on Read:", err)
			return
		}
		bytesData := buf[:n]
		stringData := string(bytesData)
		fmt.Println("Server get data:", stringData)

		// Send back data to client
		n, err = c.Write(bytesData)
		if err != nil {
			fmt.Println("Error on Write:", err)
		}
	}
}

func main() {
	socketPath := "socket.socket"
	_, err := os.Stat(socketPath)
	if err == nil {
		fmt.Println("Deleting existing socket:", socketPath)

		err = os.Remove(socketPath)
		if err != nil {
			fmt.Println("Error on removing:", err)
			return
		}
	}

	// Start Listener server.
	fmt.Println("Server is listening...")
	l, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("Listen error:", err)
		return
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			fmt.Println("Error on accept:", err)
			return
		}

		go echo(fd)
	}

}
