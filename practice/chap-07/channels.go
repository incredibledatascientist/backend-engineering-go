package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Student struct {
	Roll    int
	Name    string
	Subject string
}

func Consumer(ctx context.Context, ch <-chan Student) {
	i := 0
	for {
		select {
		case std, ok := <-ch:
			if !ok {
				fmt.Println("Error while reading, returning...")
				return
			}
			i++
			fmt.Println(i, ":- data recieved...", std)
			time.Sleep(5 * time.Second)
		}
	}
}

func Generator(ctx context.Context, ch chan<- Student) {
	// for i := 0; i <= 10; i++ {
	// 	ch <- Student{Roll: 100 + i, Name: fmt.Sprintf("Abhi-%d", i), Subject: fmt.Sprintf("Python-%d", i)}
	// 	fmt.Printf("%d :- Sleeping for %d seconds\n", i, i)
	// }
	i := 0
	for {
		select {
		case ch <- Student{Roll: 100 + i, Name: fmt.Sprintf("Abhi-%d", i), Subject: fmt.Sprintf("Python-%d", i)}:
			i++
			fmt.Printf("%d :- data sent\n", i)
			time.Sleep(1 * time.Second)

		case <-ctx.Done():
			fmt.Println("Context is cancelled now returning...")
			return
		}
	}
}

func main() {
	fmt.Println("---------- Main Start ----------")
	inputChan := make(chan Student, 1024)
	ctx, cancel := context.WithCancel(context.Background())

	// Use as goroutine for generate data
	go Generator(ctx, inputChan)

	// Use as goroutine for consume data
	go Consumer(ctx, inputChan)

	// Handle termination
	sigs := make(chan os.Signal, 1)
	fmt.Println("Waiting for the interrupt signal to exit....")
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	cancel()
	time.Sleep(5 * time.Second)
	fmt.Println("---------- Main End ----------")
}
