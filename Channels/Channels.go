package main

import "fmt"

func main() {
	// Initiates a new channel
	messages := make(chan string)

	// Calls a anonimous goroutine and sends a value to the 'message' channel.
	// There must be one send (chan <-) to each receive (<- chan).
	go func() { messages <- "ping" }()

	// Wait until 'messages' channel has some value, and then recover it.
	// There must be one receive (<- chan) to each send (chan <-).
	msg := <-messages
	fmt.Println(msg)
}
