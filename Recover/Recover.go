package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		// Prevents 'main' function to exist with the thrown panic.
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error\n", r)
		}
	}()

	mayPanic()
	fmt.Println("After mayPanic()")
}
