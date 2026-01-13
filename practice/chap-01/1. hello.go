package main

import "fmt"

var sum int
var Integer int64
var Price float64 = 10

// const sum = 10

func test() int {
	// ans := 10 / 0
	ans := 10 / sum
	fmt.Println("Ans:", ans)
	return ans
}

func main() {
	fmt.Println("Hello World.")
	sum += 1
	val := test()
	fmt.Println("Test:", val)
	fmt.Println("Int-64:", Integer+int64(val))

	out := Price / 3
	// fmt.Printf("Price: %f\n", out)
	fmt.Printf("Price: %0.8f\n", out)
}
