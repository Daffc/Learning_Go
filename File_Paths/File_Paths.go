package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// Building a filepath (portable accross systems).
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// Normalizing paths.
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))   // Drectories in filepath.
	fmt.Println("Base(p):", filepath.Base(p)) // File in filepath.

	// Checks whether a path is absolute or not.
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	// Recovering file extension from filename.
	filename := "config.json"
	ext := filepath.Ext((filename))
	fmt.Println(ext)

	// Trims extension from filename.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Recover relative path from base ("a/b") to a target ("a/b/t/file")
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	// Recover relative path from base ("a/b") to a target ("a/c/t/file")
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
