package main

import (
	"fmt"
	"sync"
)

type Countainer struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Countainer) inc(name string) {
	c.mu.Lock()
	// Ensures that mutex will be released at the end of the function/method.
	defer c.mu.Unlock() // NOTE: functions after 'defer' are stacked up to be executed at the end of the current function.
	c.counters[name]++
}

func worker(name string, n int, c *Countainer, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		c.inc(name)
	}
	wg.Done()
}

func main() {
	c := Countainer{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Initializes WaitGroup and call workers.
	wg.Add(3)
	go worker("a", 10000, &c, &wg)
	go worker("a", 10000, &c, &wg)
	go worker("b", 10000, &c, &wg)
	wg.Wait()

	fmt.Println(c.counters)
}
