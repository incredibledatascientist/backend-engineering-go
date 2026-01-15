package main

import "fmt"

var results = make(map[string]int)

func addEntry(sub string, marks int) error {
	_, ok := results[sub]
	if ok {
		return fmt.Errorf("Subject already exists.")
	}
	results[sub] = marks
	return nil
}

func readEntry(sub string) (int, error) {
	v, ok := results[sub]
	if !ok {
		return 0, fmt.Errorf("Subject '%s' not present", sub)
	}
	return v, nil
}

func readAllEntry() {
	for k, v := range results {
		fmt.Println(k, "-->", v)
	}
}

func deleteEntry(sub string) error {
	_, ok := results[sub]
	if ok {
		delete(results, sub)
		return nil
	}
	return fmt.Errorf("Subject not present in entry.")
}

func main() {
	results["html"] = 90
	results["css"] = 85
	results["js"] = 90
	results["python"] = 98
	results["go"] = 97
	// fmt.Println("Results:", results)
	readAllEntry()

	fmt.Println("-----------------")
	err := addEntry("java", 100)
	if err != nil {
		fmt.Print(err)
		return
	}
	// fmt.Println("Results:", results)
	readAllEntry()
	//  No new variable in the left side
	// v, err := ------- Only works if any v|err is new

	// err := deleteEntry("java")
	err = deleteEntry("java")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("-----------------")
	// fmt.Println("Results:", results)
	readAllEntry()
	fmt.Println("-----------------")
	val, err := readEntry("pythopn")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Result of python:", val)
}
