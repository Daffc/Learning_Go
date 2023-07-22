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

	for i := 1; i <= 5; i++ {
		// Adding a new worker to the waiting group.
		wg.Add(1)

		go func(index int) { // Setting 'i' value to the 'index', binding the correct value to each execution of the closure.
			defer wg.Done()
			worker(index)
		}(i) // Passing 'i' as value to closure in order to prevent missreading of it's value.
	}

	// Waits until al worker are finished (WaitGroup counter = 0)
	wg.Wait()
}
