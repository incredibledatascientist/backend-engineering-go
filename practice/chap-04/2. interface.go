package main

import (
	"fmt"
	"math"
)

// type Speaker interface {
// 	Speak() string
// }

// func NewSpeaker(s Speaker) {
// 	fmt.Println(s.Speak())
// }

// type Human struct {
// 	name string
// }

// func (h Human) Speak() string {
// 	return "Human is speaking..."
// }

// type Dog struct {
// 	animal bool
// }

// func (d Dog) Speak() string {
// 	return "Dog is barking..."
// }

// type Animal struct {
// 	category string
// }

// func (a Animal) Speak() string {
// 	return "Animal is making sound..."
// }

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

func ProcessShape(s Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}

// Circle
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Rectangle
type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	// // Human
	// human := Human{name: "Abhishek"}
	// NewSpeaker(human)

	// // Dog
	// dog := Dog{animal: true}
	// NewSpeaker(dog)

	// // Animal
	// animal := Animal{category: "Lion"}
	// NewSpeaker(animal)

	// -------------------------------
	// Way-1: Basic level using var
	var shape Shape
	// Circle
	shape = Circle{radius: 3}
	fmt.Println("Area of circle:", shape.Area())
	fmt.Println("Perimeter of circle:", shape.Perimeter())

	// Way-2: Production level
	// ProcessShape(c)

	// Rectangle
	shape = Rectangle{length: 10, width: 20}
	fmt.Println("Area of rectangle:", shape.Area())
	fmt.Println("Perimeter of rectangle:", shape.Perimeter())
	// ProcessShape(r)
}
