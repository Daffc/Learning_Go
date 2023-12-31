package main

import "embed"

/*
	======= HOW TO EXECUTE EXAMPLE =======

	$ mkdir -p folder
	$ echo "hello go" > folder/single_file.txt
	$ echo "123" > folder/file1.hash
	$ echo "456" > folder/file2.hash

	$ go run embed-directive.go

	======= HOW TO EXECUTE EXAMPLE =======
*/

// ==== EMBEDS VALUES IN COMPILING TIME =====

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

// ==== EMBEDS VALUES IN COMPILING TIME =====

func main() {
	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))

}
