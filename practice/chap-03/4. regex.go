package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Important functions.
	// regexp.MustCompile(pattern)
	// re.Match([]byte)

	// Regular expression
	// . - any single char

	// can not use only this without any char,
	// "*" -> not allowed, "a*" -> allowed
	// * >=0 -> occurence
	// + >=1 -> occurence
	// ? 0|1 -> occurence
	// ^ -> at strt
	// $ -> at end

	// [] -> Grouping
	// [A-Z] -> All capital letters
	// [a-z] -> All small letters
	// [0-9] -> All digits letters
	// \d -> all digits
	// \D -> non digits
	// \s -> single space char
	// \S -> non space char
	// \w -> [A-Za-z0-9_] all chars + digit + '_'
	// \W -> [A-Za-z0-9_] all chars + digit + '_', Except this

	strng := "_Abhishek"
	// regEx := "^[A-Z][a-z]*$"
	regEx := "^\\w"
	re := regexp.MustCompile(regEx)
	fmt.Println("re:", re)

	res := re.Match([]byte(strng))
	fmt.Println("res:", res)
}
