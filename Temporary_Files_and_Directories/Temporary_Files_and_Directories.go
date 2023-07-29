package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Creates file in the temporary OS folder (/tmp)
	f, err := os.CreateTemp("", "sample")
	check(err)
	defer os.Remove(f.Name())
	fmt.Println("Temp file name:", f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// Creates a file in the temporary OS folder (/tmp)
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	defer os.RemoveAll(dname)
	fmt.Println("Temp dir name:", dname)

	// Creates and writing into a new tile inside the temporary directory.
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
