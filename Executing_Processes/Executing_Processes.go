package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// Recover path of command binary
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Defining string slice with command and arguments.
	args := []string{"ls", "-a", "-l", "-h"}

	// Recovering environment variables.
	env := os.Environ()

	// Syscall to execute command, replacing go process by the command process.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
