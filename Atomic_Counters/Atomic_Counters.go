package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func worker(ops *uint64, wg *sync.WaitGroup) {
	for c := 0; c < 1000; c++ {
		// Atomic adding value (only one goroutine have access to 'ops' at time)
		atomic.AddUint64(ops, 1)
	}
	wg.Done()
}

func main() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go worker(&ops, &wg)
	}

	wg.Wait()
	fmt.Println("ops:", ops)
}
