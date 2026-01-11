package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("No values provied.")
		return
	}

	min := 0
	max := 0
	for i := 1; i < len(arguments); i++ {
		value, err := strconv.Atoi(arguments[i])
		if err != nil {
			fmt.Println("Invalid number:", arguments[i])
			continue
		}
		if i == 1 {
			min = value
			max = value
		}

		if value < min {
			min = value
		}

		if value > max {
			max = value
		}

	}

	fmt.Println("Minimum value is:", min)
	fmt.Println("Maximum value is:", max)
}
