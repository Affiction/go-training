package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var counter uint64

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for i := 0; i < 1000; i++ {
				atomic.AddUint64(&counter, 1)
				// counter += 1
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
