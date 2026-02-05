package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func defaultBuiltins(s *Shell) map[string]CommandFunc {
	return map[string]CommandFunc{
		"exit": exitCmd,
		"echo": echoCmd,
		"type": s.typeCmd,
	}
}

func exitCmd(args []string) error {
	if len(args) == 0 {
		os.Exit(0)
	}

	status, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "exit: numeric argument required")
		os.Exit(1)
	}

	os.Exit(status)
	return nil
}

func echoCmd(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

func (s *Shell) typeCmd(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(s.out, "Received no args")
		return nil
	}

	name := args[0]

	if _, ok := s.commands[name]; ok {
		fmt.Fprintln(s.out, name+" is a shell builtin")
		return nil
	}

	if path, err := exec.LookPath(name); err == nil {
		fmt.Fprintln(s.out, name+" is "+path)
		return nil
	}

	fmt.Fprintln(s.out, name+" not found")
	return nil
}
