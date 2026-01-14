package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func random(min, max, seed int) int {
	if seed > 0 {
		source := rand.NewSource(int64(seed))
		r := rand.New(source)
		return r.Intn(max-min) + min
	}
	return rand.Intn(max-min) + min
}

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		log.Fatal("Usage: min max size seed")
	}

	min, err := strconv.Atoi(arguments[1])
	if err != nil {
		log.Fatal("Invalid number:", arguments[1])
	}

	max, err := strconv.Atoi(arguments[2])
	if err != nil {
		log.Fatal("Invalid number:", arguments[2])
	}

	size, err := strconv.Atoi(arguments[3])
	if err != nil {
		log.Fatal("Invalid number:", arguments[3])
	}

	seed, err := strconv.Atoi(arguments[4])
	if err != nil {
		log.Fatal("Invalid number:", arguments[4])
	}

	for i := 0; i <= size; i++ {
		fmt.Print(random(min, max, seed), "  ")
	}

}
