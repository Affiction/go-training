package main

import "fmt"

func zeroValue(val int) {
	val = 0
}

func zeroPointer(val *int) {
	*val = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroValue(i)
	fmt.Println("zeroval:", i)

	zeroPointer(&i)
	fmt.Println("zeroptr:", i)

	// memory reference
	fmt.Println("pointer:", &i)
}
