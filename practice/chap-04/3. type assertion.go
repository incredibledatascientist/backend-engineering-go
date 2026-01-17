package main

import "fmt"

// interface{} == any (any is just a alias for empty interface.)
// func Input() interface{} {
func Input() any {
	return 201
	// return "1.2"
	// return inp{value: "1"} // unknown type
}

func main() {
	// Type Assertion
	input := Input()

	// technique-1
	v, ok := input.(int)
	if !ok {
		fmt.Println(input, " is not a integer.")
		return
	}

	// fmt.Println(input, "is a valid integer:", v)
	if ok {
		fmt.Println("Assertion successfull -> v:", v)
	}

	// _ = input.(bool) // Panic err
	// inp, ok := input.(bool)
	// if !ok {
	// 	fmt.Println("inp:", inp)
	// 	return
	// }

	// fmt.Println("in:", inp)
	// fmt.Println("Panic due to input is not bool")

	switch input.(type) {
	case int:
		fmt.Println(input, " is an int type.")
	case float64:
		fmt.Println(input, " is an float64 type.")
	case string:
		fmt.Println(input, " is an string type.")
	default:
		fmt.Println(input, " is unknown type.")
	}

}
