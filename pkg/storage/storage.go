package storage

import (
	"slices"

	"github.com/nobbmaestro/tmux-tether/pkg/tmux"
)

type SessionStorage struct {
	path     string
	Sessions []tmux.Session `yaml:"sessions"`
}

func New(path string) *SessionStorage {
	return &SessionStorage{
		path: path,
	}
}

func (s *SessionStorage) Add(sess tmux.Session) {
	if !slices.Contains(s.Sessions, sess) {
		s.Sessions = append(s.Sessions, sess)
	}
}

func (s *SessionStorage) Remove(sess tmux.Session) {
	if slices.Contains(s.Sessions, sess) {
		idx := slices.Index(s.Sessions, sess)
		s.Sessions = slices.Delete(s.Sessions, idx, idx+1)
	}
}
