package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " = ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Applying different number of arguments to variadic function
	sum(1, 2)
	sum(1, 2, 3)

	// Using slice as input to array arguments.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
