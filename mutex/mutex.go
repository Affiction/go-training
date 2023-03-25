package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex // guards
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[name]++
}

func main() {
	var wg sync.WaitGroup
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(5)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)
	go doIncrement("b", 10000)
	go doIncrement("b", 10000)

	wg.Wait()

	fmt.Println(c.counters)
}
