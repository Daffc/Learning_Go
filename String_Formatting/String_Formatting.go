package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)  // Prints structure value.
	fmt.Printf("struct2: %+v\n", p) // Prints values and struct valriables names.
	fmt.Printf("struct3: %#v\n", p) // Prints name of the package of struct definition, struct type, struct variable names and values.

	fmt.Printf("type: %T\n", p)    // Prints type of variable.
	fmt.Printf("bool: %t\n", true) // Prints boolean value.
	fmt.Printf("int: %d\n", 123)   // Prints integer value in base-10.
	fmt.Printf("bin: %b\n", 14)    // Prints value in base-2 (binary).
	fmt.Printf("char: %c\n", 33)   // Prints character corresponding to the value.
	fmt.Printf("hex: %x\n", 456)   // Prints hexadecimal representation of value.

	fmt.Printf("float1: %f\n", 78.9)        // Prints float value.
	fmt.Printf("float2: %e\n", 123400000.0) // Prints float in scientific notation.
	fmt.Printf("float3: %E\n", 123400000.0) // Prints float in scientific notation.

	fmt.Printf("str1: %s\n", "\"string\"") // Prints string value.
	fmt.Printf("str2: %q\n", "\"string\"") // Prints string as in the source value.
	fmt.Printf("str3: %x\n", "hex this")   // Prints base-16 (hexadecial) of each rune of string value.

	fmt.Printf("pointer: %p\n", &p) // Prints pointer value.

	fmt.Printf("width1: |%6d|%6d|\n", 12, 456)           // Alocates at least 6 spaces for each printeger int.
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.4555)   // Alocates at least 6 integer spaces and maximum of 2 decimal spaces for printed value.
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.4555) // Justify float values to the left.
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")        // Alocates at least 6 spaces for printed string value.
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")      // Justify string values to the left.

	s := fmt.Sprintf("sprintf: a %s", "string") // Stores formated string to string variable.
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error") // Pritns formated string to an io.Writers (Stderro, Stdout, files...)
}
