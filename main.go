package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nobbmaestro/tmux-tether/cmd"
	"github.com/nobbmaestro/tmux-tether/pkg/config"
	"github.com/nobbmaestro/tmux-tether/pkg/registry"
	"github.com/nobbmaestro/tmux-tether/pkg/service"
	"github.com/nobbmaestro/tmux-tether/pkg/storage"
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

var storePath = filepath.Join(
	os.Getenv("HOME"),
	".local",
	"state",
	"tmux-tether",
	"state.yml",
)

func run() error {
	cfg := config.ReadUserConfig(confPath)
	s := storage.New(storePath)

	if err := s.Read(); err != nil {
		return err
	}

	defer func() {
		if err := s.Write(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	ser := service.New(
		&cfg.Session,
		s,
		tmux.New(
			tmux.WithBin(cfg.TmuxCommand),
			tmux.WithHook(tmux.SessionClosedHook),
			tmux.WithHook(tmux.SessionCreatedHook),
		),
	)

	reg := registry.NewRegistry(
		registry.WithConfig(cfg),
		registry.WithConfigPath(confPath),
		registry.WithService(ser),
	)

	cmd.SetContext(reg.Context)
	cmd.SetVersionInfo(version, commit, date)

	return cmd.Execute()
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}
