package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Roll int
	Name string
}

var students []Student

func main() {
	s := Student{Roll: 101, Name: "Abhishek"}
	fmt.Printf("Type of s: %T\n", s)

	// Student struct
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	kind := t.Kind()
	fmt.Println("Type of s using reflect:", t)
	fmt.Println("Value of s using reflect:", v)
	fmt.Println("Kind:", kind)

	noOfFields := t.NumField()
	fmt.Println("No of fields:", noOfFields)

	for i := 0; i < noOfFields; i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Println("------------------------")
		fmt.Println("Field:", value)
		fmt.Println("Field Name:", field.Name)
		fmt.Println("Field Type:", field.Type)
	}

	//  slice of students
	stdType := reflect.TypeOf(students)
	stdVal := reflect.ValueOf(students)
	fmt.Println("Type of s using reflect:", stdType)
	fmt.Println("Value of s using reflect:", stdVal)

	// Only applicable for struct types not for other data type.
	// fields := stdType.NumField()
	// fmt.Println("No of fields:", fields)

}
