package main

import (
	"fmt"
	"os"
)

func main() {
	shell := NewShell(os.Stdin, os.Stdout, os.Stderr)

	if err := shell.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Shell exited with error:", err)
		os.Exit(1)
	}
}
