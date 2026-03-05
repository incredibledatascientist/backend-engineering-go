package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func status(domain, port string) string {
	addr := domain + ":" + port
	timeout := 5 * time.Second

	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return fmt.Sprintf("Dest: %s - Not Reachable (%v)", addr, err)
	}

	defer conn.Close()

	return fmt.Sprintf("Source: %s -> Dest: %s - Reachable",
		conn.LocalAddr().String(),
		conn.RemoteAddr().String(),
	)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	port := "80"

	fmt.Println("Domain Health Checker")
	fmt.Println("Type domain name (example: google.com)")
	fmt.Println("Type 'stop' to exit")

	for {
		fmt.Print("Enter domain: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		domain := strings.TrimSpace(input)

		if domain == "stop" {
			fmt.Println("Stopping checker...")
			break
		}

		result := status(domain, port)
		fmt.Println(result)
	}
}

// go run .\main.go --domain=google.com
