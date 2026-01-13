package main

import (
	"fmt"
	"strings"
)

func main() {
	// 	🔹 Comparison with other format verbs
	// Verb	Use case
	// %v	Default value (safe choice)
	// %t	Boolean only
	// %d	Integers
	// %s	String
	// %f	Floating point
	// %c	Character (rune)
	// %T	Type of value

	// f := fmt.Printf // fmt.Printf() -> f is asias for it
	// f("Formated msg: %s", "Hi")

	fmt.Printf("Equal strings:%t\n", strings.EqualFold("Abhishek", "ABhishEK"))
	strn := "Abhishek"
	// strIndex := strings.Index(strn, "ek")
	// fmt.Println("Index of 'ek':", strIndex)

	strIndex := strings.Index(strn, "Ek")
	strIndex1 := strings.Index(strn, "ek")
	fmt.Printf("Index of 'Ek':%d & ek:%d\n", strIndex, strIndex1)

	fmt.Println("Starts with Abhi:", strings.HasPrefix(strn, "Abhi"))
	fmt.Println("Starts with ABhi:", strings.HasPrefix(strn, "ABhi"))

	fmt.Println("Ends with ek:", strings.HasSuffix(strn, "ek"))
	fmt.Println("Ends with Ek:", strings.HasSuffix(strn, "Ek"))

	// 	3️⃣ Side-by-Side Comparison
	// Feature	strings.Fields()	strings.Split()
	// Split by	Any whitespace	Exact delimiter
	// Multiple spaces	✅ handled	❌ creates empty
	// Trims spaces	✅ yes	❌ no
	// Empty strings	❌ removed	✅ preserved
	// Best for	Human text	Structured data

	// Fields - Whitespace = space, tab, newline

	intro := "My   Name \nis      Abhishek kumar"
	intro2 := "My,Name,is,Abhishek,kumar"
	fields := strings.Fields(intro)
	splits := strings.Split(intro, " ")
	fmt.Println("Original string:", intro)
	fmt.Println("fields:", fields)
	fmt.Println("splits:", splits)

	fields2 := strings.Fields(intro2)
	splits2 := strings.Split(intro2, ",")
	fmt.Println("Original string2:", intro2)
	fmt.Println("fields:", fields2)
	fmt.Println("splits:", splits2)

	fmt.Println("fields2:", len(fields2))
	fmt.Println("fields2:", len(splits2))
}
