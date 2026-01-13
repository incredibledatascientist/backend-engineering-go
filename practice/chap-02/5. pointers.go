package main

import "fmt"

func increament(x int) {
	for i := 1; i <= 10; i++ {
		x++
	}
}

func increamentPointer(x *int) {
	for i := 1; i <= 10; i++ {
		(*x)++
	}
}

func main() {
	// x := 100
	// p := &x
	// fmt.Println("x =", x)
	// fmt.Println("&x =", &x)
	// fmt.Println("p =", p)
	// fmt.Println("*p =", *p)
	// fmt.Println("&p =", &p)

	// x = 200
	// fmt.Println("*p =", *p)
	// fmt.Println("&p =", &p)

	// *p = 300
	// fmt.Println("*p =", *p)
	// fmt.Println("&p =", &p)

	// ---------------------------
	x := 0
	fmt.Println("original x=", x)
	increament(x)
	fmt.Println("after increament x=", x)

	fmt.Println("original x=", x)
	increamentPointer(&x)
	fmt.Println("after increament x=", x)
}
