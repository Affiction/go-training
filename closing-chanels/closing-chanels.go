package main

import "fmt"

func main() {
	jobIds := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			id, more := <-jobIds

			if more {
				fmt.Println("Read job", id)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	go func() {

		for i := 0; i <= 3; i++ {
			jobIds <- i
			fmt.Println("Sent job", i)
		}
		close(jobIds)
	}()

	fmt.Println("Before done")
	<-done
	fmt.Println("After done")
}
