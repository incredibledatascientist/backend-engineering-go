package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Roll    int    `json:"roll"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
}

func main() {
	std := Student{101, "Abhishek Kumar", "Python"}
	fmt.Println("std:", std)

	// Marshal: go struct to json string.
	stdBytes, err := json.Marshal(&std)
	if err != nil {
		panic(err)
	}
	jsonStr := string(stdBytes)
	fmt.Println("Json String:", jsonStr)

	// Unmarshaling
	var student *Student
	err = json.Unmarshal(stdBytes, &student)
	if err != nil {
		panic(err)
	}

	fmt.Println("Struct obj:", student)
}
