package tmux

import (
	"github.com/nobbmaestro/tmux-tether/pkg/shell"
)

type Tmux struct {
	bin   string
	shell *shell.Shell
}

func New(s *shell.Shell, bin string) *Tmux {
	if bin == "" {
		bin = "tmux"
	}
	return &Tmux{
		bin:   bin,
		shell: s,
	}
}

func (t *Tmux) CreateOrSwitch(s Session) error {
	if exists := t.HasSession(s); !exists {
		if err := t.NewSession(s); err != nil {
			return err
		}
	}
	return t.SwitchClient(s)
}

func (t *Tmux) NewSession(s Session) error {
	return t.shell.CmdWithoutOutput(
		t.bin,
		"new-session",
		"-ds", s.Name,
		"-c", s.Path,
	)
}

func (t *Tmux) SwitchClient(s Session) error {
	return t.shell.CmdWithoutOutput(
		t.bin,
		"switch-client",
		"-t", s.Name,
	)
}

func (t *Tmux) HasSession(s Session) bool {
	err := t.shell.CmdWithoutOutput(
		t.bin,
		"has-session",
		"-t", s.Name,
	)
	return err == nil
}
