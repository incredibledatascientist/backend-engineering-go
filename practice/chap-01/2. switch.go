package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Args:", os.Args)
	// if len(os.Args) < 2 {
	// 	fmt.Println("Please provide value.")
	// 	return
	// }

	// argument := os.Args[1]

	// num, err := strconv.Atoi(argument)
	// if err != nil {
	// 	fmt.Println("Invalid digit string")
	// 	return
	// }

	// fmt.Println("Integer num:", num)

	// switch num {
	// case 1:
	// 	fmt.Println("1 value")
	// case 2:
	// 	fmt.Println("2 value")
	// case 3:
	// 	fmt.Println("3 value")
	// 	fallthrough
	// default:
	// 	fmt.Println("Other option")
	// }

	// switch {
	// case num < 0:
	// 	fmt.Println("Negative number")
	// case num == 0:
	// 	fmt.Println("Zero")
	// case num > 0:
	// 	fmt.Println("Positive number")
	// 	fallthrough
	// default:
	// 	fmt.Println("Other option")
	// }

	// for loop
	// for i := 0; i < 100; i++ {
	// 	fmt.Println("value of i:", i)
	// }

	// i := 0
	// for i < 10 {
	// 	fmt.Println("value of i:", i)
	// 	i++
	// }

	// Infinite loop
	// i := 0

	// for {
	// 	i++
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println("i:", i)
	// 	time.Sleep(time.Second)
	// 	if i == 10 {
	// 		break
	// 	}
	// }

	// list := []int{6, 9, 1, 2, 3, 4, 5}
	// for i, v := range list {
	// 	fmt.Printf("%d : %v\n", i, v)
	// }

	// Taking user input
	// 1. scan: Use this when you want simple values (int, float, string without spaces).
	// single input and multiple inputs
	// var name string
	// var age int
	// fmt.Println("Enter your name & age:")
	// fmt.Scan(&name, &age)
	// fmt.Printf("Your name is:%s & age is:%d\n", name, age)
	// input accept new line seprated or space : Abhishek 28, Abhishek [enter] 28

	// doesn't handle space.
	// var name string
	// fmt.Println("Enter your name:")
	// fmt.Scanln(&name)
	// fmt.Println("Your name is:", name)

	// Best Way.. : Reads data until first delimeter & including delimeter
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your full name:")
	name, err := reader.ReadString('\n')
	// name, err := reader.ReadString(' ')
	if err != nil {
		fmt.Println("Error while reading:", err)
	}

	// fmt.Printf("Your full name is:%s& size: %d\n", name, len(name))
	// fmt.Printf("Your full name is:%s& size: %d\n", name, len(name))

	clean_name := strings.TrimSpace(name)
	fmt.Printf("Clean name :%s& size: %d\n", clean_name, len(clean_name))

	// 	Enter your full name:
	//        Abhishek Kumar Suryawanshi
	// Your full name is:       Abhishek Kumar Suryawanshi
	// & size: 37
	// Clean name :Abhishek Kumar Suryawanshi& size: 26

	// Best practice with bufio.Reader — it’s what real Go devs use.
	// Error while reading: EOF , handle

	// generate a colorful image of this and explain in the tutorial and dont make too cuts just leave it.
	// os.Args : if run then 'Args: [C:\Users\ashwa\AppData\Local\Temp\go-build101184631\b001\exe\2. switch.exe]'
	// Args: [C:\Users\ashwa\OneDrive\Desktop\Go-Lang\zero-to-mastery\2. switch.exe]
	// if build then given file name. check once
	// 	The first command-line argument stored in the os.Args slice is always the name of
	// the executable. If you are using go run, you will get a temporary name and path

	// Build application: app.exe is importatnt in windows not required for linux.
	// go build file.go - created .exe windows.
	// go build -o bin/name.exe app.go
	// .\bin\name.exe
	// fmt.Println("Waiting for key press")
	// fmt.Scanln()
}
