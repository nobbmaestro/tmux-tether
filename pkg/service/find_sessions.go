package service

import (
	"os"
	"path/filepath"

	configpath "github.com/nobbmaestro/tmux-tether/pkg/config/parsers"
	"github.com/nobbmaestro/tmux-tether/pkg/tmux"
)

func isSessionDir(dir string, markers []string) bool {
	for _, indicator := range markers {
		if _, err := os.Stat(filepath.Join(dir, indicator)); err != nil {
			return false
		}
	}

	return true
}

func findSessionsInDir(
	dir string,
	markers []string,
	currentDepth int,
	maxDepth int,
) ([]tmux.Session, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var sessions []tmux.Session

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		entryDir := filepath.Join(dir, entry.Name())

		if isSessionDir(entryDir, markers) {
			sessions = append(sessions, tmux.Session{
				Name: entry.Name(),
				Path: configpath.New(entryDir),
			})
		}

		if maxDepth >= 0 && currentDepth >= maxDepth {
			continue
		}

		nestedSessions, err := findSessionsInDir(
			entryDir,
			markers,
			currentDepth+1,
			maxDepth,
		)
		if err != nil {
			return nil, err
		}

		sessions = append(sessions, nestedSessions...)
	}

	return sessions, nil
}

func (s *Service) FindSessions() ([]tmux.Session, error) {
	var sessions []tmux.Session

	for _, dir := range s.cfg.Dirs {
		found, err := findSessionsInDir(dir.String(), s.cfg.Markers, 0, s.cfg.Depth)
		if err != nil {
			return nil, err
		}

		sessions = append(sessions, found...)
	}

	return sessions, nil
}
