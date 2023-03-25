package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add1(a, b int) int {
	return a + b
}

func add2(a, b int) (sum int) {
	sum = a + b
	return
}

func main() {

	res := add(2, 2)
	fmt.Println(res)
	fmt.Println(add1(2, 2))
	fmt.Println(add2(2, 2))
}
