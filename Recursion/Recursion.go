package main

import "fmt"

// Recursive factorial Function.
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Applying recursions over a Closure.
	var fib func(n int) int // <- this explicit declaration is needed
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
