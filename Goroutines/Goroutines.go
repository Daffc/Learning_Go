package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	// Calling a goroutine
	go f("goroutines")

	// Calling a goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Letting 1 second pass until the process exits.
	time.Sleep(time.Second)
	fmt.Println("done")
}
