package main

import "fmt"

func init() {
	fmt.Println("FROM INIT")
}

func main() {

	fmt.Println("Basic form")
	if 15%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if i := 12; i%2 == 0 {
		fmt.Println("12 % 2")
	}

	if false {
		fmt.Println("false block")
	} else if true {
		fmt.Println("true block")
	} else {
		fmt.Println("else block")
	}
}
