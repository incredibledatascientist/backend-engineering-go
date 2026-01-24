package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// for i := 1; i <= 10; i++ {
	// 	ch <- i
	// }
	ch <- 100
	// close(ch)
}

func main() {
	fmt.Println("---------------- Main Start ----------------")
	channel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go producer(channel, &wg)

	fmt.Println("Wait for the goroutines to complete...")
	fmt.Println("Read<-", <-channel)
	wg.Wait()
	fmt.Println("Goroutines completed")

	fmt.Println("---------------- Main End ----------------")
}
