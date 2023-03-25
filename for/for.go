package main

import "fmt"

const defaultForIterationNumber = 5

func main() {

	fmt.Println("Short version")
	i := 0
	for i < defaultForIterationNumber {
		fmt.Println(i)
		// i++
		// i = i + 1
		i += 1
	}

	fmt.Println("Full version")
	for i := 0; i < defaultForIterationNumber; i++ {
		fmt.Println(i)
	}

	fmt.Println("Break")
	for i := 0; i < defaultForIterationNumber; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("Continue")
	for i := 0; i < defaultForIterationNumber; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			continue
		}
	}
}
