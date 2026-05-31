package shell

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

type Shell struct {
}

func New() *Shell {
	return &Shell{}
}

func (s *Shell) Cmd(cmd string, args ...string) (string, error) {
	var stdout bytes.Buffer

	command := exec.Command(cmd, args...)
	command.Stdin = os.Stdin
	command.Stdout = &stdout

	if err := command.Run(); err != nil {
		return "", err
	}

	return strings.TrimSuffix(stdout.String(), "\n"), nil
}

func (s *Shell) CmdWithoutOutput(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout

	return command.Run()
}

func (s *Shell) CmdPassthrough(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.Run()
}
