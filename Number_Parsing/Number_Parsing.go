package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Parsing from string to float64 (string, bits)
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// Parsing from string to Int64 (string, base, bits)
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// Parsing from hexvalue representative strig to Int64
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// Parsing from string to Uint64
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Parsing string to base-10 int.
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Parsing error example.
	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}
