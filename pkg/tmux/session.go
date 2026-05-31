package tmux

import (
	"os"
	"path/filepath"

	configpath "github.com/nobbmaestro/tmux-tether/pkg/config/parsers"
)

type Session struct {
	Name string
	Path string
}

func isSessionDir(dir string, indicators []string) bool {
	for _, indicator := range indicators {
		if _, err := os.Stat(filepath.Join(dir, indicator)); err != nil {
			return false
		}
	}

	return true
}

func findSessionsInDir(
	dir string,
	indicators []string,
	currentDepth int,
	maxDepth int,
) ([]Session, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var sessions []Session

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		entryDir := filepath.Join(dir, entry.Name())

		if isSessionDir(entryDir, indicators) {
			sessions = append(sessions, Session{
				Name: entry.Name(),
				Path: entryDir,
			})
		}

		if maxDepth >= 0 && currentDepth >= maxDepth {
			continue
		}

		nestedSessions, err := findSessionsInDir(
			entryDir,
			indicators,
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

func FindSessions(
	searchDirs []configpath.Path,
	sessionIndicators []string,
	searchDepth int,
) ([]Session, error) {
	var sessions []Session

	for _, dir := range searchDirs {
		found, err := findSessionsInDir(dir.String(), sessionIndicators, 0, searchDepth)
		if err != nil {
			return nil, err
		}

		sessions = append(sessions, found...)
	}

	return sessions, nil
}
