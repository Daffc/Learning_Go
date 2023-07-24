package main

import "os"

func main() {

	// Tries to creta a file without permission.
	_, err := os.Create("/file")
	if err != nil {
		// Throus error and stop execution.
		panic(err)
	}
}
