package main

import "fmt"

// Function to be tested/benchmarked;
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(IntMin(1, 2))
}
