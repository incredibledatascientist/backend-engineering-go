package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	// var num1, num2 float64
	// Setup logging.

	// Ensure dir exists
	logDir := "logs"
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		// log.Fatal: use when the program cannot continue due to unrecoverable runtime issues
		// such as permission errors, missing directories/files, or invalid configuration.
		log.Fatal("Failed to create directory:", err)
	}

	// Ensure log file exist
	logFile := "app.log"
	logFilePath := filepath.Join(logDir, logFile)
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	defer f.Close()

	logger := log.New(f, "cal ", log.LstdFlags)
	logger.Println("calculator application started.")

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: calculator <number> <operator> <number>")
		fmt.Println("Operations: +, -, *, /")
		return
	} else if len(arguments) != 4 {
		logger.Print("Error: invalid number of arguments:", arguments[1:])
		fmt.Println("Error: invalid number of arguments")
		fmt.Println("Expected: <number> <operator> <number>")
		return
	}

	num1, err := strconv.ParseFloat(arguments[1], 64)
	if err != nil {
		logger.Println("Invalid number:", err)
		fmt.Println("Invalid number:", err)
		return
	}

	num2, err := strconv.ParseFloat(arguments[3], 64)
	if err != nil {
		logger.Println("Invalid number:", err)
		fmt.Println("Invalid number:", err)
		return
	}

	operator := arguments[2]
	switch operator {
	case "+":
		fmt.Printf("Addition of %f+%f=%f", num1, num2, num1+num2)
	case "-":
		fmt.Printf("Substraction of %f-%f=%f", num1, num2, num1-num2)
	case "*":
		fmt.Printf("Multiplication of %f*%f=%f", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			logger.Println("Error: division by zero is not allowed")
			fmt.Println("Error: division by zero is not allowed")
			return
		}
		fmt.Printf("Division of %f/%f=%f", num1, num2, num1/num2)

	default:
		logger.Printf("Error: invalid operator '%s'", operator)
		fmt.Printf("Error: invalid operator '%s'\n", operator)
		fmt.Println("Hint! Allowed operators: + - * /")
	}
}
