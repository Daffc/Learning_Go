package main

import (
	"fmt"
	"time"
)

func main() {

	// Initiates a buffered channel with jobs.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Initiates a ticker that will fire every 200 milliseconds.
	limiter := time.Tick(200 * time.Millisecond)

	// Loops over requests, executing then every fire of the ticker.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Defining a buffer of tickers.
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Closure gouroutine that will set one value in ticker buffer every 200 milliseconds.
	go func() {
		// NOTE: In this implementation, this approach does not perform as desired as the loops are "buffered" every 200 milliseconds,
		// meaning that for longer waits, 'burstyLimiter' will receive valuer in intervals less than 200 milliseconds.
		// A better approach would be to use an infinite loop with a time.Sleep() between sets, that which only one thick would be "buffered".
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Initiates a new buffered channel of jobs
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	fmt.Println("-------------------")
	// Loops trought all jobs.
	for req := range burstyRequests {
		<-burstyLimiter // Blocks next job to be executed until there are a value in 'busrtLimiter' to be received.
		fmt.Println("request", req, time.Now())
	}
}
