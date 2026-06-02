package git

import "github.com/nobbmaestro/tmux-tether/pkg/shell"

type Option func(*Git)

type Git struct {
	bin   string
	shell *shell.Shell
}

func New(opts ...Option) *Git {
	g := &Git{
		bin:   "git",
		shell: &shell.Shell{},
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

func WithShell(s *shell.Shell) Option {
	return func(t *Git) {
		t.shell = s
	}
}

func WithBin(bin string) Option {
	return func(t *Git) {
		if bin != "" {
			t.bin = bin
		}
	}
}

func (g *Git) Clone(url string) error {
	return g.shell.CmdPassthrough(
		g.bin,
		"clone",
		url,
	)
}
