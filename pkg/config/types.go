package config

import path "github.com/nobbmaestro/tmux-tether/pkg/config/parsers"

type SessionConfig struct {
	// Directories to search for potential tmux sessions
	Dirs []path.Path `yaml:"dirs"`

	// Maximum depth to search below each directory in SearchDirs.
	// 0 = only the directory itself, 1 = one level deep, -1 = unlimited recursion.
	Depth int `yaml:"depth"`

	// Files or directories that marks a valid session root
	Markers []string `yaml:"markers"`

	// Directory names or glob patterns to exclude from scanning.
	Exclude []string `yaml:"exclude"`
}

type UserConfig struct {
	// Session configs
	Session SessionConfig `yaml:"session"`

}
