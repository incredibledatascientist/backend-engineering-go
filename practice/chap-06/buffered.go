package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	file, err := os.Open("input.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	time_start := time.Now()
	reader := bufio.NewReader(file)
	buf := make([]byte, 10)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		read := string(buf[:n])
		fmt.Print(read)
	}
	duration := time.Since(time_start)
	fmt.Println("duration:", duration)
}
