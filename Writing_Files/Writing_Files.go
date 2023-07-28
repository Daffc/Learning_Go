package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Dumping strings / bytes to a file.
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// Creating a new file for writing
	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close() // Schedule file close

	// Writing bytes to file.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2) // Returns the number of bytes written.
	check(err)
	fmt.Printf("write %d bytes\n", n2)

	// Writting string to file.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Flushes writes to the file.
	f.Sync()

	// Buffers Writes to the file.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("Write %d bytes\n", n4)

	// Flushes al buffered writes.
	w.Flush()
}
