package service

import (
	"github.com/nobbmaestro/tmux-tether/pkg/config"
	"github.com/nobbmaestro/tmux-tether/pkg/tmux"
)

type Service struct {
	cfg  *config.SessionConfig
	tmux *tmux.Tmux
}

func New(cfg *config.SessionConfig, t *tmux.Tmux) *Service {
	return &Service{
		tmux: t,
		cfg:  cfg,
	}
}
