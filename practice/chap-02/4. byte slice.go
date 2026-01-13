package main

import "fmt"

func main() {
	// b := []byte("abhishek")
	b := make([]byte, 12)
	b = []byte("Byte slice €")

	fmt.Println("byte slice: ", b)
	fmt.Println("byte slice size: ", len(b))
	fmt.Printf("byte string : %s\n", b)
	fmt.Println("byte string :", string(b))

	r := '€'
	fmt.Println("Rune data type:", r)

	// Note :In Go, len(string) returns the number of BYTES, not the number of characters.
	unicode := "€"
	fmt.Println("unicode string:", unicode)
	fmt.Println("unicode size:", len(unicode))
	fmt.Println("unicode bytes:", []byte(unicode))
	fmt.Println("unicode bytes size:", len([]byte(unicode)))

	// len(string) -> only for string it return number of bytes not for array/slice or other.
	lst := []string{"a", "b", "€"}
	fmt.Println("lst string:", lst)
	fmt.Println("lst size:", len(lst))

	// Deleting an element from a slice
	// Technique-1:- Split into two slice and merge
	oldSlice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("oldSlice:", oldSlice)
	// sliceA := oldSlice[:3]
	// sliceB := oldSlice[4:]
	// newSlice := oldSlice[:3] + oldSlice[4:]
	// fmt.Println("sliceA:", sliceA)
	// fmt.Println("sliceB:", sliceB)
	// newSlice := append(sliceA, sliceB...)
	// fmt.Println("newSlice:", newSlice)

	// Technique-2:-
	oldSlice[3] = oldSlice[len(oldSlice)-1]
	newSlice := oldSlice[:len(oldSlice)-1]
	fmt.Println("newSlice tech-2:", newSlice)

	// 	buf := make([]byte, 1024)
	// n, _ := file.Read(buf)
}
