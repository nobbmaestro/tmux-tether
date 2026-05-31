package config

import path "github.com/nobbmaestro/tmux-tether/pkg/config/parsers"

func GetDefaultUserConfig() *UserConfig {
	return &UserConfig{
		TmuxCommand: "",
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
		Picker: PickerConfig{
			Pointer: " ",
			Prompt:  "  ",
		},
	}
}
