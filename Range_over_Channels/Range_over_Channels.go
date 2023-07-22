package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Recoverig (receiving) entryies of "queue" channel buffer using range
	for elem := range queue {
		fmt.Println(elem)
	}

	// After loop, buffered channel "queue" is empty.
}
