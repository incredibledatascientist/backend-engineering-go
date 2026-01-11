package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type File struct {
	name string
	mode string
}

func FindBinaryFile(path, file string) File {
	pathSplit := filepath.SplitList(path)
	for _, dir := range pathSplit {
		// fmt.Println("-->", dir)
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
				// Is it executable?
				if mode&0111 != 0 {
					// fmt.Println("Yes it is executable & path:", fullPath)
					return File{name: file, mode: "binary"}
				}

				// fmt.Println("Regular file")
				return File{name: file, mode: "regular"}

			}

		}
	}

	return File{}
}

func main() {
	fmt.Println("----------------- Welcome to which utility application -------------")

	// Searches for the executable only in directories listed in the PATH environment variable.
	// Only reads path present in envirment variable.
	// path := os.Getenv("PATH")
	path := "C:\\Users\\ashwa\\OneDrive\\Desktop\\Go-Lang\\zero-to-mastery\\applications"
	// fmt.Println("path:", path)

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an input")
		return
	}

	for _, file := range arguments[1:] {
		// fmt.Println("file:", file)
		result := FindBinaryFile(path, file)
		fmt.Println("result:", result)
	}
}
