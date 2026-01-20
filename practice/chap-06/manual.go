package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 10)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		read := string(buf[:n])
		fmt.Print(read)
	}
}
