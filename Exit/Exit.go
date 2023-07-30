package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!")
	os.Exit(3) // Exists program (without execuing stacked difer's) with code 3.
}
