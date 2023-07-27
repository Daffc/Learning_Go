package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// Applys regular expressions over strings and returns if string is valid.
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Defines a variable with the regex expressions struct (more Efficient than the inline applied expressions as in the example above).
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))                                        // Checks if the strings machs the regex struct.
	fmt.Println(r.FindString("patch peach"))                                   // Finds the firstsubstring that maches de regex struct.
	fmt.Println("idx", r.FindStringIndex("peach  punch"))                      // Recover slice with the positions of the first and the last rune of the first match substring.
	fmt.Println(r.FindStringSubmatch("peach punch"))                           // Recover a slice containing match strings and submatches of regex in the match string.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))                      // Recover a slice containing the indexes of the match strings and submatches of regex in the match string.
	fmt.Println(r.FindAllString("peach punch pinch", -1))                      // Recover a slice containing all substrings that match the regex.
	fmt.Println("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1)) // Recover a slice containing all the indexes of the match strings and submatches of regex in the match string.
	fmt.Println(r.FindAllString("peach punch pinch", 2))                       // Recovers a slice with the first n substrings that matches the regex.
	fmt.Println(r.Match([]byte("peach")))                                      // Checks it the []byte value matches the regular expressions.

	// Definies a variales with a regex expressions that will panics instead of returning error.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // Replaces substring that matches the regex second string argument.

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper) // Applies string functions to match substring.
	fmt.Println(string(out))

}
