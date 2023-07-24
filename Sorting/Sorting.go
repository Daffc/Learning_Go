package main

import (
	"fmt"
	"sort"
)

func main() {

	// Sorting slice of strings.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// Sorting slice of integers.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// Chacks if a slice of integers is sorted.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
