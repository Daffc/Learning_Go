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
	// creates a directory with '755' permission.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Schedules exclusion of direcoty.
	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// Creates all files of the described in path.
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	c, err := os.ReadDir(("subdir/parent"))
	check(err)

	// Lists all the content inside 'subdir/parent' directory.
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir()) // Checks if entry is a directory.
	}

	// Changes working directory (similar to 'cd' command).
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Getts content in the corrent directory.
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir(("../../.."))
	check(err)

	// Recusivelly visits every file and directory inside target element.
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fmt.Println(" ", p, info.IsDir())
	return nil
}
