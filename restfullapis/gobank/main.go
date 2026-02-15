package main

import (
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

	// Create table
	err = store.createGobankTable()
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(cfg, store)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
