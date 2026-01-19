package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("------------ main start ------------")

	processID := os.Getpid()
	fmt.Println("process id:", processID)
	sigs := make(chan os.Signal, 1)
	// Windows: only SIGINT works via Ctrl+C
	// signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT, syscall.SIGUSR1) // SIGUSR1 not supported in windows
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println("<- waiting for signal...")
	sig := <-sigs
	fmt.Println("Signal recieved:", sig)
	switch sig {
	case syscall.SIGINT:
		fmt.Println("Intererrupt signal")
	case syscall.SIGTERM:
		fmt.Println("Termination signal")
	// can't handle this
	// case syscall.SIGKILL:
	// 	fmt.Println("Kill signal")
	default:
		fmt.Println("Un-known signal.")
	}

	fmt.Println("------------ main End ------------")
}
