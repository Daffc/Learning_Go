package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// The 'devfault' optin prevents select to wait for the 'send' and 'receive' events from channels.
	// If the cannel action is not available, the code preceeds with the execution (non-blocking)
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received sinal", sig)
	default:
		fmt.Println("no activity")
	}

}
