package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Student struct {
	Roll    int
	Name    string
	Subject string
}

func Consumer(ctx context.Context, ch <-chan Student, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for std := range ch {
		i++
		fmt.Println(i, ":- data recieved...", std)
		time.Sleep(5 * time.Second)
	}
	fmt.Println("All data read successfully now returning...")
}

func Generator(ctx context.Context, ch chan<- Student, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		close(ch) // Closing a chan
	}()
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

	var wg sync.WaitGroup

	// Use as goroutine for generate data
	wg.Add(1)
	go Generator(ctx, inputChan, &wg)

	// Use as goroutine for consume data
	wg.Add(1)
	go Consumer(ctx, inputChan, &wg)

	// Handle termination
	sigs := make(chan os.Signal, 1)
	fmt.Println("Waiting for the interrupt signal ....")
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	cancel()
	fmt.Println("Waiting for the goroutine to return....")
	wg.Wait()
	fmt.Println("---------- Main End ----------")
}
