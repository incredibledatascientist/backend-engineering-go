package main

import (
	"log"
	"os"
)

func main() {
	// Open/Create file in append mode
	logfile := "app.log"
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Panic("File open err:", err)
	}
	// Setup logging.
	logger := log.New(f, "app ", log.LstdFlags)

	arguments := os.Args
	if len(arguments) == 1 {
		logger.Printf("No arguments provide for: %s", os.Args[0])
		return
	}

	input := os.Args[1]
	logger.Print("input provided:", input)

	logger.Print("----------Thank you for providing log file----------")
}
