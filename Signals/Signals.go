package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Defines a buffered channel of os.Signal.
	sigs := make(chan os.Signal, 1)

	// Register expected signals (SIGINT and SIGTERM) to channel 'sigs'.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Defines a buffered channel of os.Signal, locking process end.
	done := make(chan bool, 1)

	// Defines goroutine to await for a signal (SIGINT or SIGTERM).
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaitng singal")
	<-done
	fmt.Println("exiting")
}
