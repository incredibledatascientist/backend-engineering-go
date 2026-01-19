// -----------------------------------------------
// stdin, stdout, stderr (UNIX + Go) — Short Notes
// -----------------------------------------------
//
// Every program starts with 3 open files:
//
// stdin  (fd 0) -> input  (keyboard, file, pipe)
// stdout (fd 1) -> normal output
// stderr (fd 2) -> error output
//
// UNIX uses numbers 0, 1, 2 internally (file descriptors).
//
// Go exposes these safely as:
//   os.Stdin   -> standard input
//   os.Stdout  -> standard output
//   os.Stderr  -> standard error
//
// Why this is useful:
// - Programs can be piped together (echo "hi" | app)
// - Output and errors can be separated
// - Input can come from keyboard, file, or another program
//
// Proof (shell examples):
//   go run app.go > out.txt        // redirect stdout
//   go run app.go 2> err.txt       // redirect stderr
//   go run app.go > all.txt 2>&1   // redirect both
//   echo data | go run app.go      // pipe into stdin
//
// Best practice in Go:
// - Print normal output to os.Stdout
// - Print errors/logs to os.Stderr
// -----------------------------------------------

package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Fprint(os.Stdout, "Enter your name:")
	fmt.Fscan(os.Stdin, &name)

	fmt.Fprintf(os.Stdout, "\nYour name is: %s", name)
	fmt.Fprintln(os.Stderr, "\nThis is just a err.")
}
