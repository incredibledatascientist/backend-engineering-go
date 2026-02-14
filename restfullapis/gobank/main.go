package main

import (
	"log"
	"time"
)

func main() {
	server := NewAPIServer(
		"localhost:8080",
		10*time.Second,
		10*time.Second,
		60*time.Second,
	)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
