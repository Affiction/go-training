package main

import "fmt"

func fact(i int) int {
	if i == 0 {
		return 1
	}

	return i * fact(i-1)
}

func main() {
	var fib func(n int) int

	fib = func(i int) int {
		if i < 2 {
			return i
		}

		return fib(i-1) + fib(i-2)
	}

	fmt.Println(fact(10))
	fmt.Println(fib(7))
}
