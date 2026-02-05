package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type CommandFunc func(args []string) error

type Shell struct {
	in       io.Reader
	out      io.Writer
	err      io.Writer
	reader   *bufio.Reader
	commands map[string]CommandFunc
}

func NewShell(in io.Reader, out, err io.Writer) *Shell {
	s := &Shell{
		in:     in,
		out:    out,
		err:    err,
		reader: bufio.NewReader(in),
	}

	s.commands = defaultBuiltins(s)

	return s
}

func (s *Shell) Run() error {
	for {
		fmt.Fprint(s.out, "OhmyGosh$ ")

		line, err := s.reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(s.err, "Error reading command:", err)
			return err
		}

		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		fields := strings.Fields(trimmed)
		if len(fields) == 0 {
			continue
		}

		name := fields[0]
		args := fields[1:]

		if cmdFn, ok := s.commands[name]; ok {
			if err := cmdFn(args); err != nil {
				fmt.Fprintln(s.err, "Error:", err)
			}
			continue
		}

		if err := runExternal(name, args, s.in, s.out, s.err); err != nil {
			fmt.Fprintln(s.err, "Error:", err)
		}
	}
}
