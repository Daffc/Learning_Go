package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Defining new set of flags for subcommand 'foo'
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// Defining new set of flags for subcommand 'bar'
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Checks if there are at lead 2 arguments (program call and subcommand)
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Checks first argument (subcommand)
	switch os.Args[1] {
	case "foo": // Parses 'foo' arguments.
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())

	case "bar": // Parses 'bar' arguments.
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  leve:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())

	default: // If there is not a valid subcommand.
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

}
