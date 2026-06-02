package main

import (
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

func main() {
	cfg := config.ReadUserConfig(confPath)
	store := storage.New(storePath)

	err := store.Read()
	if err != nil {
		os.Exit(2)
	}

	defer func() {
		if cerr := store.Write(); err == nil {
			err = cerr
		}
	}()

	ser := service.New(
		&cfg.Session,
		store,
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

	err = cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
