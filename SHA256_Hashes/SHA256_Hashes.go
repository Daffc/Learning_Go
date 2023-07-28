package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"
	h := sha256.New() // Defines a new Hash functio.

	// Hashing string 's'
	h.Write([]byte(s))

	// Recuvering hash result.
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
