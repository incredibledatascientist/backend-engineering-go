package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(st, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for i := st; i <= end; i++ {
		fmt.Print(i, "=")
	}
	fmt.Println()
	time.Sleep(100 * time.Microsecond)
}

func main() {
	fmt.Println("---- goroutines ----")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go printMessage(i, 5, &wg)
	}

	wg.Wait()

	// time.Sleep(10 * time.Second)
}

//	func myPrint(start, finish int) {
//		for i := start; i <= finish; i++ {
//			fmt.Print(i, " ")
//		}
//		fmt.Println()
//		// time.Sleep(100 * time.Microsecond)
//	}
// func myPrint(start, finish int) {
// 	for i := start; i <= finish; i++ {
// 		fmt.Printf("[G%d:%d] ", start, i)
// 	}
// 	fmt.Println()
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		go myPrint(i, 5)
// 	}
// 	time.Sleep(time.Second)
// }
