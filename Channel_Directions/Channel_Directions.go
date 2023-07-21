package main

import "fmt"

// Receives 'msg' and sends it to the 'pings' channel.
// Specify that 'pings' is only for sends.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Await (receive) message from 'ping' and transfer it to 'pongs'.
// Specifyes that 'pings' is a channel only for receives and 'pongs' only fir  sends.
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)

	// Await (receive) until 'pongs' has a value.
	fmt.Println(<-pongs)
}
