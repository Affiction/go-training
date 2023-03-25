package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			worker(i)
		}(i)
	}

	wg.Wait()
}

// func main() {
// 	var wg sync.WaitGroup
// 	done := make(chan struct{})
// 	wq := make(chan interface{})
// 	workerCount := 2

// 	for i := 0; i < workerCount; i++ {
// 		wg.Add(1)
// 		go doit(i, wq, done, &wg)
// 	}

// 	for i := 0; i < workerCount; i++ {
// 		wq <- i
// 	}

// 	close(done)
// 	wg.Wait()
// 	fmt.Println("all done!")
// }

// func doit(workerId int, wq <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
// 	fmt.Printf("[%v] is running\n", workerId)
// 	defer wg.Done()
// 	for {
// 		select {
// 		case m := <-wq:
// 			fmt.Printf("[%v] m => %v\n", workerId, m)
// 		case <-done:
// 			fmt.Printf("[%v] is done\n", workerId)
// 			return
// 		}
// 	}
// }
