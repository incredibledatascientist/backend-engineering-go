package main

func main() {
	// sl := []byte("ab")
	// ml := []byte("AB")
	// // r := "👉abhi" // string not rune
	// // r := "👉" // string not rune
	// r := '👉' // rune
	// fmt.Println(sl, string(sl))
	// fmt.Println(ml, string(ml))
	// fmt.Printf("byte string: %s\n", sl)
	// fmt.Printf("byte string: %s\n", ml)
	// fmt.Println("rune bytes: ", r)
	// fmt.Printf("rune as character %c\n", r) // only expect single rune char "// r := '👉'"

	// Convert Runes to text
	// fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	// iterate through the string.
	// for _, ch := range sl {
	// 	fmt.Println(ch)
	// }

	// // iterate through the sring rune not able to do for single rune char
	// for _, ch := range r {
	// 	// fmt.Println(ch)
	// 	fmt.Println(string(ch))
	// }

	// The length of the string is the same as the number of characters found in the string
	// fmt.Println(len(sl))

	// 	usually not true for byte slices because Unicode characters usually require
	// more than one byte.
	// fmt.Println(len([]byte(sl)))

	// // Print an existing string as characters
	// aString := "abhishek"
	// for _, v := range aString {
	// 	fmt.Printf("%c", v)
	// }
	// fmt.Println()

	// Print an existing string as runes
	// for _, v := range aString {
	// 	fmt.Printf("%x ", v)
	// }

	// Converting from int to string
	// fmt.Println(string(65)) // convert into single char

	// st := strconv.Itoa(65)
	// fmt.Println(st)

	// str := strconv.FormatInt(int64(65), 10)
	// fmt.Print(str)

	// The strings package

}
