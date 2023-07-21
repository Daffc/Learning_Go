package main

import "fmt"

func main() {
	// Initiates a new channel buffer with 2 positions.
	messages := make(chan string, 2)

	// Pushes two messages to the buffer
	messages <- "buffered"
	messages <- "channel"

	// Pull messages from chanel in FIFO order.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
