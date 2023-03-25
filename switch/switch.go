package main

import (
	"fmt"
	"time"
)

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type is integer")
	case string:
		fmt.Println("Type is string")
	default:
		fmt.Printf("Type is %T\n", i)
	}
}

func main() {
	fmt.Println("Basic")
	i := 1
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Default")
	}

	fmt.Println("Multiple")
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Workday")
	}

	fmt.Println("Empty switch")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	default:
		fmt.Println("Evening")
	}

	x := 2e10
	printType(1)
	printType(3.14)
	printType(true)
	printType("Hello")
	printType(x)
}
