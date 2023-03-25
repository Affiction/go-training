package main

import (
	"fmt"
	"time"
)

func worker(c chan bool) {
	fmt.Print("Working...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
	c <- true
}

func main() {
	done := make(chan bool)

	go worker(done)

	<-done
}
