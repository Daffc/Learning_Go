package main

import (
	"fmt"
	"sort"
)

// Defining a new type 'byLength' to describe a slice os strings.
type byLength []string

// Redefining Methods of sorting for 'byLength' type.
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool { // Criteria for ordering elements.
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
