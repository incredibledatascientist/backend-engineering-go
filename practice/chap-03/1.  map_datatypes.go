package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println("map:", m)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println("map:", m)
	fmt.Println("map size:", len(m))

	// delete element from map
	delete(m, "b")

	fmt.Println("after delete map:", m)
	fmt.Println("after delete map size:", len(m))

	// Iterate through map
	for k, v := range m {
		fmt.Println(k, "-->", v)
	}

	// check key exists or not
	value, ok := m["a"]
	// value, ok := m["b"]
	if !ok {
		fmt.Println("Key not available")
		// return
	}
	fmt.Println("value is:", value)

	// 	In real-world applications, if a function accepts a map argument,
	// then it should check that the map is not nil before working with it.
	// m = nil                              // after nil if you access key it will crash
	// fmt.Println("after nil, a:", m["a"]) // this will give zeroed value but assigning will crash

	// m["k"] = 10 // try to add in nil map will crash.
	if m != nil {
		fmt.Println("get a value:", m["a"])
	}
}
