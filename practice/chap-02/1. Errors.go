package main

import "fmt"

func division(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divisor can't be 0.")
	}
	result := a / b
	fmt.Print("")
	return float64(result), nil
}

func main() {
	res, err := division(10, 3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Result is:", res)
}
