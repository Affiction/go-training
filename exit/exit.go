package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Exit Code - 3")
	defer fmt.Println("deferred message")

	os.Exit(3)
}