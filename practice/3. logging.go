package main

import (
	"fmt"
	"log"
	"os"
)

// log.Fatal() :- Logs the message and exit(1) not run defer.
// log.Fatal() :- Logs the message and call Panic() but run defer before exit.
func main() {
	// fmt.Println("------ main start --------")
	// defer fmt.Println("main program completed")
	// fmt.Println("start-1")
	// arguments := os.Args
	// if len(arguments) == 1 {
	// 	log.Fatal("No arguments provided!")
	// }

	// fmt.Println("start-2")
	// log.Panic("Argument provided.")

	// fmt.Println("------ main ends --------")

	// ****** Writing to a custom log file **********

	// LOGFILE := filepath.Join(os.TempDir(), "temp_log.log")
	LOGFILE := "C:\\Users\\ashwa\\OneDrive\\Desktop\\Go-Lang\\zero-to-mastery\\logger.log"

	fmt.Println("logfile:", LOGFILE)
	// f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("File open err:", err)
		return
	}

	defer f.Close()
	// Flags:
	// Quick comparison table (cheat sheet)
	// Flag	Adds
	// Ldate	Date
	// Ltime	Time
	// Lmicroseconds	Microseconds
	// Lshortfile	File + line
	// Llongfile	Full path + line
	// LUTC	UTC time
	// Lmsgprefix	Prefix after timestamp
	// LstdFlags	Date + time

	// logger := log.New(f, "prefix ", log.LstdFlags)

	logger := log.New(f, "prefix ", log.Lshortfile)
	logger.SetFlags(log.LstdFlags | log.Lshortfile)

	logger.Println("First log message.")
	logger.Println("Second log message.")
}
