package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// Worker awaits until there is a job to be received.
	for j := range jobs {
		fmt.Println("Worker", id, "started job", j, "...")
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "finished job.", j, ".")
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Initiates all workers.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Making all jobs available to the workers through the channel 'jobs'
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Receinving all jobs results (in the order they are send by the workers).
	for a := 1; a <= numJobs; a++ {
		fmt.Println("result", a, ":", <-results)
	}
}
