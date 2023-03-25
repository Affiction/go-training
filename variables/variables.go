package main

import "fmt"

func main() {
	var a = "Hello"
	fmt.Println("infer type", a)

	var b int = 1
	fmt.Println("with type declaration", b)

	// Multiple variables declarations
	var c, d = 1, 2
	fmt.Println("Multiple variables declarations", c, d)

	c, d = d, c
	fmt.Println("Change order of multiple declaration", c, d)

	var e int
	fmt.Println("Without corresponding initialization", e)

	f :=  3.14
	fmt.Println("Short var declaration", f)
}
