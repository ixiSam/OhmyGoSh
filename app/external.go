package main

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func runExternal(name string, args []string, in io.Reader, out, err io.Writer) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = in
	cmd.Stdout = out
	cmd.Stderr = err

	if errRun := cmd.Run(); errRun != nil {
		var ee *exec.Error
		if errors.As(errRun, &ee) && ee.Err == exec.ErrNotFound {
			fmt.Fprintln(out, name+": not found")
			return nil
		}

		if _, ok := errRun.(*exec.ExitError); ok {
			// Command executed but exited with non-zero status; REPL should keep running.
			return nil
		}

		return fmt.Errorf("%s: execution failed: %v", name, errRun)
	}

	return nil
}
