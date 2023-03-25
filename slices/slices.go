package main

import "fmt"

func main() {
	a := make([]string, 5)
	fmt.Println(a, len(a), cap(a))

	b := make([]int, 5, 10)
	fmt.Println(b, len(b), cap(b))
	b = append(b, 1, 2, 3)
	fmt.Println(b, len(b), cap(b))
	
	// Copies
	c:= make([]int, len(b), cap(b))
	copy(c, b)
	// c := b[:]
	b[len(b)-1] = 1000
	fmt.Println(b, len(b), cap(b))
	fmt.Println(c, len(c), cap(c))
}
