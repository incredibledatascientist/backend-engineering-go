// Packages group related code (functions, variables, constants).
// Only package main can create a runnable program.
// Other packages are libraries, not executables.

// Does NOT apply to package names
// func PrintData() {}   // public
// func printData() {}   // private

// old way 1.16
// go get github.com/spf13/cobra

// Modern way (Go ≥ 1.16):
// go install github.com/spf13/cobra@latest

// go mod init appname [Initialize mod]
// go mod tidy [Update mod]
// mod.sum [version tracking dont update manually/ dont touch]

// Folder = Package
// main = entry point
// Uppercase = public
// go.mod = dependency list
// go.sum = dependency proof
// go mod tidy = clean & download

package main

import (
	"application/services"
	"application/test"
	"fmt"
)

func init() {
	fmt.Println("This is init() from main package.")
}

//	func init() {
//		fmt.Println("This is init() from main package. 2")
//	}
func main() {
	fmt.Println("============== Main program starts ================")
	fmt.Println(services.Greet())
	test.TestGreet()
}
