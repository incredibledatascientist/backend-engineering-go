package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// Client is connecting
	socketPath := "socket.socket"
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	for {
		buffer := bufio.NewReader(os.Stdin)
		text, _ := buffer.ReadString('\n')
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Write error:", err)
			break
		}

		buf := make([]byte, 128)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read err: %+v, %d", err, n)
			return
		}

		bytesData := buf[:n]
		strData := string(bytesData)
		fmt.Println("Read :", strData)

		if strings.TrimSpace(strData) == "STOP" {
			fmt.Println("Exiting the unix client!")
			return
		}

		time.Sleep(2 * time.Second)
	}
}
