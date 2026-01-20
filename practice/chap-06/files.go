package main

import (
	"fmt"
	"io"
	"os"
)

// func lineByLine(file string) error {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	reader := bufio.NewReader(f)

// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			fmt.Println("Some error while reading...")
// 			break
// 		}

// 		fmt.Printf("[%q]", line)

// 	}

// 	return nil
// }

func main() {
	fmt.Println("------------- main start -----------")
	// err := lineByLine("input.log")
	// if err != nil {
	// 	fmt.Println("Error :", err)
	// }

	// // Read entire file into memory
	// // Use for Config files, Small JSON/YAML, Environment files
	// dataBytes, err := os.ReadFile("input.log")
	// if err != nil {
	// 	panic(err)
	// }
	// dataStr := string(dataBytes)
	// fmt.Println("complete data read:-")
	// fmt.Printf("%q", dataStr)

	// Step 2: Manual buffered reading
	file, err := os.Open("input.log")
	defer file.Close()

	if err != nil {
		panic(err)
	}
	buf := make([]byte, 1024)
	count := 0
	for {
		count++
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			panic("Err")
		}
		fmt.Print(string(buf[:n]))
	}

	println()
	println("Total syscal:", count)
	fmt.Println("------------- main end -----------")
}
