package main

import "fmt"

// The return values of a function can be named
func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
		return min, max
	}
	min = x
	max = y
	return
}

func genID() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {

	// func() {
	// 	fmt.Println("Anonymous functions...")
	// }()

	// func(name string) {
	// 	fmt.Println("Hello ", name)
	// }("Abhishek")

	// res := func(name string) string {
	// 	return fmt.Sprintf("Hello %s...", name)
	// }("Abhishek")

	// fmt.Println(res)

	// Lamda / annonymous functions.
	greet := func(name string) string {
		return "Hello " + name
	}
	fmt.Println(greet("Abhishek"))

	// The return values of a function can be named
	min, max := minMax(3, 2)
	fmt.Printf("Min: %d & Max:%d\n", min, max)

	// Closures
	// It will return func() int -> function not a directly int
	// fmt.Println(generate()) // Prints memory address not id
	generate := genID()
	fmt.Println(generate())
	fmt.Println(generate())
	fmt.Println(generate())
}
