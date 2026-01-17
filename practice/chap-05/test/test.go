package test

import "fmt"

func init() {
	fmt.Println("This is init() from test package.")
}

func TestGreet() {
	fmt.Println("Test package: Hello...")
}
