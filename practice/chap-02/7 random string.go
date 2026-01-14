package main

import (
	"fmt"
	"math/rand"
)

const (
	MIN = 0
	MAX = 94
)

func randomNum(min, max int) int {
	return rand.Intn(max-min) + min
}

func getAlphaNum(n int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

func getString(len int64) string {
	// 	'!' = 33   ← first printable character
	// '~' = 126  ← last printable character
	// random(0 to 93) + 33

	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := randomNum(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func main() {

	strng := ""
	for i := 0; i < 10; i++ {
		// 127 excluded → max 126
		r := randomNum(33, 127)
		strng += string(r)
	}
	fmt.Println("random string:", strng)
	fmt.Println("getStr:", getString(5))
	fmt.Println("getAlphaNum:", getAlphaNum(8))
}
