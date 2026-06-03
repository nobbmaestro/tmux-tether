package tmux

import (
	"fmt"

	"github.com/nobbmaestro/tmux-tether/pkg/shell"
)

type Option func(*Tmux)

type Tmux struct {
	bin   string
	shell *shell.Shell
}

func New(opts ...Option) *Tmux {
	t := &Tmux{
		bin:   "tmux",
		shell: &shell.Shell{},
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

func WithShell(s *shell.Shell) Option {
	return func(t *Tmux) {
		t.shell = s
	}
}

func WithBin(bin string) Option {
	return func(t *Tmux) {
		if bin != "" {
			t.bin = bin
		}
	}
}

func WithHook(h *Hook) Option {
	return func(t *Tmux) {
		err := t.SetHook(h)
		if err != nil {
			fmt.Println("Setting hook failed!")
		}
	}
}

func (t *Tmux) NewSession(s Session) error {
	return t.shell.CmdWithoutOutput(
		t.bin,
		"new-session",
		"-ds", s.Name,
		"-c", s.Path.String(),
	)
}

func (t *Tmux) SwitchClient(s Session) error {
	return t.shell.CmdWithoutOutput(
		t.bin,
		"switch-client",
		"-t", s.Name,
	)
}

func (t *Tmux) SwitchLastSession() error {
	return t.shell.CmdWithoutOutput(
		t.bin,
		"switch-client",
		"-l",
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

func (t *Tmux) SetHook(h *Hook) error {
	return t.shell.CmdWithoutOutput(
		t.bin,
		"set-hook",
		"-g",
		string(h.Type),
		string(h.Cmd),
	)
}
