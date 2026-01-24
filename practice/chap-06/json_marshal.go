package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Roll int `json:"roll,omitempty"` // If Roll = 0 then don't add in JSON.
	// Roll     int    `json:"roll"` // If Roll = 0 then don't add in JSON.
	Name     string `json:"name"`
	Subject  string `json:"subject"`
	Username string `json:"-"`
}

var students []Student

func main() {
	// std := Student{101, "Abhishek Kumar", "Python", "abhishek"}
	std := Student{Name: "Abhishek", Subject: "Python", Username: "abhishek@supercloudnow.com"}
	std2 := Student{Roll: 101, Name: "Neelu", Subject: "Python", Username: "neelu@supercloudnow.com"}
	fmt.Println("Student:", std)
	students = append(students, std)
	students = append(students, std2)

	// Marshal: go struct to json string.
	// stdBytes, err := json.Marshal(&std)
	jsonStudent, err := json.Marshal(&students)
	if err != nil {
		panic(err)
	}

	jsonStr := string(jsonStudent)
	fmt.Println("Json String:", jsonStr)

	// Unmarshaling
	var goStudent []Student
	err = json.Unmarshal(jsonStudent, &goStudent)
	if err != nil {
		panic(err)
	}

	fmt.Println("Struct obj:", goStudent)
	// fmt.Printf("jsonStudent type:%T:", jsonStudent) // []bytes
	fmt.Printf("jsonStr type:%T:\n", jsonStr)
	fmt.Printf("goStudent type:%T\n", goStudent)

	// Encoding & Decoding work on streaming of data.
	// 	// Serialize serializes a slice with JSON records
	// func Serialize(e *json.Encoder, slice interface{}) error {
	//  return e.Encode(slice)
	// }

	// 	// DeSerialize decodes a serialized slice with JSON records
	// func DeSerialize(e *json.Decoder, slice interface{}) error {
	//  return e.Decode(slice)
	// }

	// decoder := json.NewDecoder(r.Body)
	// decoder.DisallowUnknownFields()

	// err := decoder.Decode(&u)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	fmt.Println("------------------ Encoding/Decoding --------------------")
	fmt.Println("students:", students)
}
