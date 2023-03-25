package main

import (
	"fmt"
	"time"
)

func basicLimiter() {
	requests := make(chan int, 5)

	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	limitter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limitter
		fmt.Println("request", req, time.Now())
	}
}

// In this implementation, we "burst" the first initial requests(limited by 3)
func burstyLimiter() {
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

func main() {
	println("Basic Limiter")
	basicLimiter()

	println("Bursty Limiter")
	burstyLimiter()
}
