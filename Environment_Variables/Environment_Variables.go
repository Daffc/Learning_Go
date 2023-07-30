package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Setting environment variable FOO (only for the program execution scope).
	os.Setenv("FOO", "1")

	// Getting envirnment variales values.
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	// Recovers all environment variables and prints their names.
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
