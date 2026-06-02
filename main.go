package main

import (
	"os"
	"path/filepath"

	"github.com/nobbmaestro/tmux-tether/cmd"
	"github.com/nobbmaestro/tmux-tether/pkg/config"
	"github.com/nobbmaestro/tmux-tether/pkg/registry"
	"github.com/nobbmaestro/tmux-tether/pkg/service"
	"github.com/nobbmaestro/tmux-tether/pkg/tmux"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

var confPath = filepath.Join(
	os.Getenv("HOME"),
	".config",
	"tmux-tether",
	"tmux-tether.yml",
)

func main() {
	cfg := config.ReadUserConfig(confPath)

	ser := service.New(
		&cfg.Session,
		tmux.New(
			tmux.WithBin(cfg.TmuxCommand),
		),
	)

	reg := registry.NewRegistry(
		registry.WithConfig(cfg),
		registry.WithConfigPath(confPath),
		registry.WithService(ser),
	)

	cmd.SetContext(reg.Context)
	cmd.SetVersionInfo(version, commit, date)

	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
