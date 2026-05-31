package git

import "github.com/nobbmaestro/tmux-tether/pkg/shell"

type Git struct {
	bin   string
	shell *shell.Shell
}

func New(s *shell.Shell, bin string) *Git {
	if bin == "" {
		bin = "git"
	}
	return &Git{
		bin:   bin,
		shell: s,
	}
}

func (g *Git) Clone(url string) error {
	return g.shell.CmdPassthrough(
		g.bin,
		"clone",
		url,
	)
}
