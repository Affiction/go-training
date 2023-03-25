package main

import (
	"fmt"
	"math"
)

const myConst = "myConst"

func main() {
	fmt.Println(myConst)

	const bigNumber = 5_000_000
	fmt.Println(bigNumber)

	const a = 3e200
	fmt.Println(a)

	fmt.Println(math.Sin(bigNumber))
}
