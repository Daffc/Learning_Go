package main

import (
	"fmt"
	"time"
)

// Defining 'worker' fuction that will be user by a goroutine and will send a value to 'done' when it's process if finished.
func worker(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	// Awaits until a value is send to 'done'
	<-done
}
