package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // Keep receiving jobs until channel is closed ('more' is false).
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true // Indicates in channel 'done' that all jobs have been processed
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	close(jobs) // Closes Channel, indicting to receives (in the second value) that the channel is closed
	fmt.Println("send all jobs")

	<-done
}
