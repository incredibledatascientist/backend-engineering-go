package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.Create("input.log") // Open file in only read mode
	// file, err := os.Create("input.log") // Open file in only write mode
	file, err := os.OpenFile("inputNew.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// If file not exists create, else append
	// defer file.Close() // Never close file before error check
	if err != nil {
		panic(err)
	}

	defer file.Close() // Never close file before error check
	for i := 11; i <= 20; i++ {
		msg := fmt.Sprintf("input-log %d\n", i)
		file.Write([]byte(msg))
	}
}
