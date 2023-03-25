package main

import (
	"fmt"
	"time"
)

func main() {
	messageChan := make(chan string)

	go func() {
		time.Sleep(time.Second)
		messageChan <- "Hello"
	}()

	msg := <-messageChan

	fmt.Println(msg)
}
