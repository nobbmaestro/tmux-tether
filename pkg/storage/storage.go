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

func sameName(sess tmux.Session) func(tmux.Session) bool {
	return func(e tmux.Session) bool { return e.Name == sess.Name }
}

func (s *SessionStorage) Add(sess tmux.Session) {
	if !slices.ContainsFunc(s.Sessions, sameName(sess)) {
		s.Sessions = append(s.Sessions, sess)
	}
}

func (s *SessionStorage) Remove(sess tmux.Session) {
	idx := slices.IndexFunc(s.Sessions, sameName(sess))
	if idx != -1 {
		s.Sessions = slices.Delete(s.Sessions, idx, idx+1)
	}
}
