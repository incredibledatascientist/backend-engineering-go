package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// fmt.Println("ab", " - ", []byte("ab"))
	// fmt.Println("Pradeep", " - ", []byte("Pradeep"))
	fmt.Println([]byte("7"))
	x := 7
	fmt.Println(unsafe.Sizeof(x))
}

// 010101
// 4 -> 8421
// 7 -> 0111 (binary)
// 8 bit -> 0000 0111

// bit -> 00 01 10 11 | -> 0 | 1

// bytes

// 1 bit -> 1 or 0

// 1 bytes = 8 bit

// 10 -> 0 1 2 3...9
// binary -> 0 1

// 1 byte = 2^1 = 2 0|1
// 8 byte = 2^1 = 2 0|1
// 8 bit = 2^8 0 to 155 256 | 2^32

// 32 bit = 0 to 255.8.8.8
// 32 bit = 255.255.255.255

// Final Clear Answer

// Byte = Raw 8-bit data.
// ASCII = Rulebook that assigns characters to certain byte values.

// Byte is storage.
// ASCII is interpretation.
