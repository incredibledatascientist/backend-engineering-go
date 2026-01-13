package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// r := rand.New(rand.NewSource(1))

	// Generate same number every time
	source := rand.NewSource(5)
	random := rand.New(source)
	fmt.Println("random:", random.Intn(3))     // 0 to 2 (3 excluded)
	fmt.Println("random:", random.Intn(10))    // 0 to 9 (10 excluded)
	fmt.Println("random:", random.Intn(10)+10) // 10 to 19
}
