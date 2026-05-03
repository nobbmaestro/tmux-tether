package main

import (
	"os"

	"github.com/nobbmaestro/tmux-tether/cmd"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	cmd.SetVersionInfo(
		version,
		commit,
		date,
	)

	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
