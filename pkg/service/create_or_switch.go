package service

import "github.com/nobbmaestro/tmux-tether/pkg/tmux"

func (s *Service) CreateOrSwitch(sess tmux.Session) error {
	if exists := s.tmux.HasSession(sess); !exists {
		if err := s.tmux.NewSession(sess); err != nil {
			return err
		}
	}
	return s.tmux.SwitchClient(sess)
}
