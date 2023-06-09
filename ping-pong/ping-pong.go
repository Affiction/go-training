package main

import "fmt"

func ping(c chan<- string, msg string) {
	c <- msg
}

func pong(ping <-chan string, pong chan<- string) {
	msg := <-ping
	pong <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "my ping")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
