package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	// Receiving 2 return values.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Ingnoring first value
	_, c := vals()
	fmt.Println(c)
}
