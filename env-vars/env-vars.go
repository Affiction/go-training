package main

import (
	"fmt"
	"os"
)

func main() {
	p := fmt.Println
	os.Setenv("A", "test")

	p("A", os.Getenv("A"))
	p("B", os.Getenv("B"))
}