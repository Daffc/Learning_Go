package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	// Recover "index", "value" from Array/Slice
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}

	// Iterating over keys and values of a map.
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Iterating over each key of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Iterating over each character (rune) of a string (i = index, c = value)
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
