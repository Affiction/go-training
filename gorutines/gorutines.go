package main

import (
	"fmt"
	"time"
)

func myFunc(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
	}
}

func main() {

	myFunc("First")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	
	go myFunc("Second")
	go myFunc("Third")

	time.Sleep(time.Second)

	fmt.Println("done")
}
