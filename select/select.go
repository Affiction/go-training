package main

import (
	"fmt"
	"time"
)

// func main() {
// 	chen1 := make(chan string)
// 	chen2 := make(chan string)

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		chen1 <- "hello"
// 	}()

// 	fmt.Println("After first")

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		chen2 <- "world"
// 	}()

// 	fmt.Println("After second")

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-chen1:
// 			fmt.Println(msg1)
// 		case msg2 := <-chen2:
// 			fmt.Println(msg2)
// 		}
// 	}

// 	fmt.Println("After for with select")
// }

func write(c chan<- string) {
	msg := "some message"
	// time.Sleep(2 * time.Second)
	c <- msg
}

func read(c <-chan string) {
	for {
		select {
		case msg := <-c:
			fmt.Println("from select:", msg)
			return 
			default: 
			fmt.Println("from default")
		}
	}
}

func main() {
	messages := make(chan string)

	println("Before first goroutine")
	go read(messages)

	println("After first goroutine")
	go write(messages)

	time.Sleep(2 * time.Second)
}
