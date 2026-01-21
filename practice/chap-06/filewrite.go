package main

import (
	"os"
)

func main() {
	// // file, err := os.Create("write.txt") // Oncre for create
	// file, err := os.OpenFile("write.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	//
	// buf := make([]byte, 1024)
	// data := []byte("This is new write in my file.\n")
	// data := "This is new write in my file."
	// fmt.Fprintf(file, string(data))

	// user WriteString
	// n, err := file.WriteString(string(data))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%d bytes written.\n", n)

	// // Using bufio -> Mandatory to flush
	// writer := bufio.NewWriter(file)
	// // n, err := writer.WriteString(string(data))
	// n, err := writer.WriteString("I love coding.\n")
	// if err != nil {
	// 	panic(err)
	// }
	// writer.Flush()
	// fmt.Printf("%d bytes written.\n", n)

	// // Using io.WriteString()
	// n, err := io.WriteString(file, string("Write data using io.WriteString"))
	// if err != nil {
	// 	panic(nil)
	// }
	// fmt.Printf("%d bytes written.\n", n)

	// // using Write([]byte)
	// n, err := file.Write([]byte("This is writting using Write().\n"))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%d bytes written.\n", n)

	// WriteFile write complete file at once for configs and small files
	err := os.WriteFile("writeFile.txt", []byte("Complete file write at once"), 0644)
	if err != nil {
		panic(err)
	}

}
