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

type PickerConfig struct {
	// Picker pointer icon
	Pointer string `yaml:"pointer"`

	// Picker prompt icon
	Prompt string `yaml:"prompt"`
}

type UserConfig struct {
	TmuxCommand string `yaml:"tmux_command"`
	GitCommand  string `yaml:"git_command"`

	// Session configs
	Session SessionConfig `yaml:"session"`

	// Picker configs
	Picker PickerConfig `yaml:"picker"`
}
