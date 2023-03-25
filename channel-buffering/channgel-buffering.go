package main

import (
	"fmt"
)

func main() {
	messageChan := make(chan string, 2)

	messageChan <- "Hello"
	messageChan <- "World"

	fmt.Println(<-messageChan)
	fmt.Println(<-messageChan)
}
