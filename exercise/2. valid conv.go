package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide atleast one input.")
		return
	}

	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, v := range arguments[1:] {
		_, err := strconv.Atoi(v)
		if err == nil {
			nInts++
			total++
			continue
		}

		_, err = strconv.ParseFloat(v, 64)
		if err == nil {
			nFloats++
			total++
			continue
		}

		invalid = append(invalid, v)
	}

	fmt.Println("#read:", total)
	fmt.Println("#nInt:", nInts)
	fmt.Println("#nFloats:", nFloats)

	if len(invalid) > total {
		fmt.Println("Too much values invalid.")
	}

}
