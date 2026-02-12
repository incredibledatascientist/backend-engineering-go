package main

import (
	"fmt"
)

func main() {

	fmt.Println("------ main golang ----------")

	server := NewAPIServer("localhost:8080")
	server.Run()
}
