package service

import (
	"github.com/nobbmaestro/tmux-tether/pkg/config"
	"github.com/nobbmaestro/tmux-tether/pkg/storage"
	"github.com/nobbmaestro/tmux-tether/pkg/tmux"
)

type Service struct {
	cfg     *config.SessionConfig
	storage *storage.SessionStorage
	tmux    *tmux.Tmux
}

func New(
	c *config.SessionConfig,
	s *storage.SessionStorage,
	t *tmux.Tmux,
) *Service {
	return &Service{
		cfg:     c,
		storage: s,
		tmux:    t,
	}
}
