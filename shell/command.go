package shell

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunCommand(command string, args ...string) string {
	cmd := exec.Command(command, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Sprintf("error: %v\nstderr: %v\n", err, stderr.String())
	}

	return stdout.String()
}

func ShellCommand(command string) string {
	split := safeSplit(command)
	l := len(split)
	if l > 1 {
		return RunCommand(split[0], split[1:]...)
	} else {
		return RunCommand(split[0])
	}
}
