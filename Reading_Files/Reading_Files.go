package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Reads file.
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// Opens file for multiple operations.
	f, err := os.Open("/tmp/dat")
	check(err)

	// Creates a buffer
	b1 := make([]byte, 5)
	n1, err := f.Read(b1) // Reads from the file to the buffer 'b1'
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1])) // Prints Retrived runes and size in bytes.

	// Moves os.File pointer to the 6th character if (6 character relative tot he origin of the file).
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// Reproduce the same behavior of the previous read implementation but with the 'ReadAtLeast' function from 'io' library.
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Setts file pointer to the beggining of the file.
	_, err = f.Seek(0, 0)
	check(err)

	// Reads from the file directtly into a buffer.
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Closing file (could be scheduled with defer).
	f.Close()
}
