package main

import (
	"log"
	"time"
)

func main() {
	cfg := ServerConfig{
		Addr:         "localhost:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	server := NewAPIServer(cfg)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
