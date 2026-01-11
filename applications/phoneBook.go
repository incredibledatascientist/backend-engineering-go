package main

import (
	"fmt"
	"os"
	"strings"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(surname string) *Entry {
	for i, v := range data {
		// if v.Surname == surname { // Case Sensitive

		// Case In-sensitive
		if strings.EqualFold(v.Surname, surname) {
			return &data[i]
		}
	}
	return nil
}

func list() {
	fmt.Println("------------- Phonebook List ----------")
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		// log.Panic("Please provide atleast one action to perform")
		fmt.Printf("Usage: %s search|list <arguments>\n", arguments[0])
		return
	}

	// Add data into phonebook app.
	data = append(data, Entry{Name: "Abhishek", Surname: "Kumar", Tel: "9399938409"})
	data = append(data, Entry{Name: "Abhi", Surname: "Dahariya", Tel: "9575682764"})
	data = append(data, Entry{Name: "Prince", Surname: "Sharma", Tel: "9876543210"})
	data = append(data, Entry{Name: "Nitesh", Surname: "Sonwani", Tel: "9876543210"})
	data = append(data, Entry{Name: "Suvendu", Surname: "Dey", Tel: "9878987898"})
	data = append(data, Entry{Name: "Pradeep", Surname: "Barik", Tel: "8081767789"})

	// Differential options
	switch arguments[1] {

	// search command
	case "search":
		if len(arguments) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println("No record found.")
			return
		}
		fmt.Println(*result)

	// list command
	case "list":
		list()

	// all other options
	default:
		fmt.Println("Invalid option.")
	}
}
