package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Some error while reading...")
			break
		}

		fmt.Printf("[%q]", line)

	}

	return nil
}

func main() {
	fmt.Println("------------- main start -----------")
	err := lineByLine("input.log")
	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Println("------------- main end -----------")
}
