package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)

	// Wait 2s
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	
	fmt.Println("after first goroutine")

	select {
	case res := <-c1:
		fmt.Println(res)
	// Wait 1s(this case will win)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	fmt.Println("after second goroutine")

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
