package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("----------------- Welcome to which utility application -------------")

	// Searches for the executable only in directories listed in the PATH environment variable.
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an input")
		return
	}

	file := arguments[1]
	fmt.Println("file:", file)

	path := os.Getenv("PATH")
	// fmt.Println("path:", path)

	pathSplit := filepath.SplitList(path)
	for _, dir := range pathSplit {
		// fmt.Println(ind, "-->", dir)
		fullPath := filepath.Join(dir, file)

		// Does it exist?
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			// fmt.Println("fullPath:", fullPath)
			// fmt.Println("fileInfo:", fileInfo)

			mode := fileInfo.Mode()
			// fmt.Println("mode:", mode)

			// Is it a regular file
			if mode.IsRegular() {
				fmt.Println("Regular file")

				// Is it executable?
				if mode&0111 != 0 {
					fmt.Println("Yes it is executable & path:", fullPath)
				}
			} else if mode.IsDir() {
				fmt.Println("It is a directory")
			}
		}
	}
}
