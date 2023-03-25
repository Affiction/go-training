package main

import "fmt"

func main() {
	a := make([]int, 5)
	b := make(map[string]int)

	fmt.Println(a, len(b))
	fmt.Println(b, len(b))

	b["test1"] = 1
	b["test2"] = 2
	fmt.Println(b, len(b))

	b["test3"] = 3
	fmt.Println(b, len(b))
	delete(b, "test2")
	fmt.Println(b, len(b))

	// Check existing value
	val, ok := b["test3"]
	fmt.Println(val, ok)

	val1, ok1 := b["test4"]
	fmt.Println(val1, ok1)
}
