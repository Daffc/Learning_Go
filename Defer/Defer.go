package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt") // Opens file
	defer closeFile(f)                // Stacks closing file to be executed at the end of the 'main' function.
	writeFile(f)                      // Writes the file
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
