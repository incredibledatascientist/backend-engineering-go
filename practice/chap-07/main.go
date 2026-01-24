package test

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Student struct {
	Roll    int
	Name    string
	Subject string
}

func Producer(ctx context.Context, ch chan<- Student, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		ch <- Student{Roll: 100 + i, Name: fmt.Sprintf("Abhi-%d", i), Subject: fmt.Sprintf("Python-%d", i)}
		fmt.Printf("Sleeping for %d seconds\n", i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}

func Consumer(ctx context.Context, ch <-chan Student) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("consumer stopped:", ctx.Err())
			return

		case std, ok := <-ch:
			if !ok {
				fmt.Println("channel closed, consumer exiting")
				return
			}
			fmt.Println("student:", std)
		}
	}
}

func main() {
	fmt.Println("---------- Main Start ----------")
	// Read write will be blocked for unbuffered channel
	// channel := make(chan Student) // channel with size 0
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Buffered channels are unblocking.
	channel := make(chan Student, 1)

	wg.Add(1)
	go Producer(ctx, channel, &wg)

	go Consumer(ctx, channel)

	fmt.Println("Waiting for all data producing...")
	wg.Wait()
	//
	sigs := make(chan bool)
	cancel()
	fmt.Println("---------- Main End ----------")
}
