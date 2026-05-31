package tmux

import "os/exec"

func CreateOrSwitch(s Session) error {
	exists, err := HasSession(s)
	if err != nil {
		return err
	}

	if !exists {
		if err := NewSession(s); err != nil {
			return err
		}
	}

	return SwitchClient(s)
}

func NewSession(s Session) error {
	cmd := exec.Command(
		"tmux",
		"new-session",
		"-ds", s.Name,
		"-c", s.Path,
	)
	return cmd.Run()
}

func SwitchClient(s Session) error {
	cmd := exec.Command(
		"tmux",
		"switch-client",
		"-t", s.Name,
	)
	return cmd.Run()
}

func HasSession(s Session) (bool, error) {
	cmd := exec.Command(
		"tmux",
		"has-session",
		"-t", s.Name,
	)

	err := cmd.Run()
	if err == nil {
		return true, nil
	}

	if exitErr, ok := err.(*exec.ExitError); ok {
		if exitErr.ExitCode() != 0 {
			return false, nil
		}
	}

	return false, err
}
