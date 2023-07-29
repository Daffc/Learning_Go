package main

import (
	"fmt"
	"os"
)

func main() {
	// Recovers program call and all arguments.
	argsWithProg := os.Args
	// Recovers only arguments.
	argsWithoutProg := os.Args[1:]
	// Recovers one agument.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
