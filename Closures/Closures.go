package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func intSub(a int) int {
	i := 0
	i -= a
	return i
}

func main() {

	// Initiates a variable "nextInt" containing "intSub" closure
	nextInt := intSeq()

	// Each time "nextInt" is called, it's "i" value updates.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Another instance of "intSub" value could be instantiated with its own scope variables.
	newInts := intSeq()
	fmt.Println(newInts())
}
