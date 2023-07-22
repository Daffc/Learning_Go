package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			// Options will be selected:
			select {
			case <-done: // When there are a value to be received in 'done'.
				return
			case t := <-ticker.C: // Every fire of the ticker.
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Wits 1600 Milliseconds ( 3 ticks and 100 milliseconds.)
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop() // Stop the ticker.
	done <- true  // Send a value to the 'done' channel (stopping ticker loop).
	fmt.Println("Ticker stopped")
}
