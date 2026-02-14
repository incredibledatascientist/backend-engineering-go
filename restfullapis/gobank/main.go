package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	cfg := ServerConfig{
		Host:         "localhost",
		Port:         8080,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("store: %+v\n", store)

	server := NewAPIServer(cfg, store)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
