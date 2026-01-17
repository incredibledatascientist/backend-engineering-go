package main

import "fmt"

func main() {
	// Benefit of using 'map[string]interface{}' is our value is in original data type
	// 'map[string]string' this will change the value to string type.
	// map[string]any
	// m := make(map[string]interface{})
	m := make(map[string]any)
	m["a"] = 100
	m["b"] = "abc"
	m["c"] = 199.5
	fmt.Println("Map m:", m)

	for k, v := range m {
		fmt.Println(k, "-->", v)

		// because v is any type/ interface{}
		// switch v.(type) {
		// case int:
		// 	fmt.Println(v, " is int")
		// case string:
		// 	fmt.Println(v, " is string")
		// case float64:
		// 	fmt.Println(v, " is float64")
		// }
	}
}
