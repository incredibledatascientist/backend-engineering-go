package services

import "fmt"

func init() {
	fmt.Println("This is init() from services package.")
}

func Greet() string {
	return "services package: Good morning..."
}

func privateGreet() string {
	return "services package: Good morning..."
}
