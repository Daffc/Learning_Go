package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Setts new scanner to the user input.
	scanner := bufio.NewScanner(os.Stdin)

	// Reads inputs from the user (separated by '\n')
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text()) // Fiters text applying upper case.
		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
