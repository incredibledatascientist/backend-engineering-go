package main

import (
	"encoding/csv"
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
	for _, v := range data {
		fmt.Println(v)
	}
}

func readFromCSV(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	// reader := csv.NewReader(f)
	reader, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	return reader, nil

}

func writeToCSV(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(f)
	// writer.Comma = rune('\t')
	// writer.Comma = '\t'
	for _, row := range data {
		record := []string{row.Tel, row.Name, row.Surname}
		writer.Write(record)
	}
	writer.Flush()
	return nil
}

func main() {
	// arguments := os.Args
	// if len(arguments) == 1 {
	// 	// log.Panic("Please provide atleast one action to perform")
	// 	fmt.Printf("Usage: %s search|list <arguments>\n", arguments[0])
	// 	return
	// }
	arguments := os.Args
	if len(arguments) == 1 {
		// log.Panic("Please provide atleast one action to perform")
		fmt.Printf("Usage: %s input output\n", arguments[0])
		return
	}

	// Add data into phonebook app.
	data = append(data, Entry{Name: "Abhishek", Surname: "Kumar", Tel: "9399938409"})
	data = append(data, Entry{Name: "Abhi", Surname: "Dahariya", Tel: "9575682764"})
	data = append(data, Entry{Name: "Prince", Surname: "Sharma", Tel: "9876543210"})
	data = append(data, Entry{Name: "Nitesh", Surname: "Sonwani", Tel: "9876543210"})
	data = append(data, Entry{Name: "Suvendu", Surname: "Dey", Tel: "9878987898"})
	data = append(data, Entry{Name: "Pradeep", Surname: "Barik", Tel: "8081767789"})

	// // Differential options
	// switch arguments[1] {

	// // search command
	// case "search":
	// 	if len(arguments) != 3 {
	// 		fmt.Println("Usage: search Surname")
	// 		return
	// 	}
	// 	result := search(arguments[2])
	// 	if result == nil {
	// 		fmt.Println("No record found.")
	// 		return
	// 	}
	// 	fmt.Println(*result)

	// // list command
	// case "list":
	// 	list()

	// // all other options
	// default:
	// 	fmt.Println("Invalid option.")
	// }

	fmt.Println("------------ Main start ----------")
	// list() // List the contacts.
	if len(arguments) == 3 {
		input := os.Args[1]
		// output := os.Args[2]
		// lines, err := readFromCSV("phonebook.csv")

		// Reading from csv file.
		lines, err := readFromCSV(input)
		if err != nil {
			fmt.Println(err)
			return
		}

		// List records
		for i, line := range lines {
			fmt.Println(i, "-->", line)
		}
		fmt.Printf("Successfully read contacts from %s csv file\n", input)

		// Writing to a csv file
		// err := writeToCSV(output)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Printf("Successfully stored contacts into %s csv file\n", output)
	}
	fmt.Println("------------ Main end ------------")
}
