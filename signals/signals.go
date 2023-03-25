package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	p := fmt.Println


	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		sig := <-sigs
		fmt.Println()
		p("-->", sig)
		done <- true
	}()

	p("awaiting signal")
	<-done
	p("exiting")
}
