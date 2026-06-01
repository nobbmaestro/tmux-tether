package config

import (
	path "github.com/nobbmaestro/tmux-tether/pkg/config/parsers"
	"github.com/nobbmaestro/tmux-tether/pkg/tmux"
)

func GetDefaultUserConfig() *UserConfig {
	return &UserConfig{
		TmuxCommand: "",
		GitCommand:  "",
		Session: SessionConfig{
			Dirs: []path.Path{
				path.New("~/repos"),
			},
			Markers: []string{
				".git",
			},
			Depth:   0,
			Exclude: []string{},
		},
		Switch: SessionSwitchConfig{
			Markers: []tmux.Session{},
		},
		Picker: PickerConfig{
			Pointer: " ",
			Prompt:  "  ",
		},
	}
}
