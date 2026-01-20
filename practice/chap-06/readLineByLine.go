package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	// Read line by line.
	// file, err := os.Open("input.log")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// reader := bufio.NewReader(file)
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Print(line)
	// }

	// Read word by word.
	file, err := os.Open("input.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	reg := regexp.MustCompile("[^\\s]")
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		// fmt.Print(line)
		// words := strings.Split(line, " ")
		words := reg.FindAllString(line, -1)
		// fmt.Print(words)
		for _, word := range words {
			fmt.Print(word)
		}
		fmt.Println()
	}
	fmt.Println()
}
