package tmux

import (
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
