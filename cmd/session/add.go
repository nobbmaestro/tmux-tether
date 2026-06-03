package session

import (
	configpath "github.com/nobbmaestro/tmux-tether/pkg/config/parsers"
	"github.com/nobbmaestro/tmux-tether/pkg/registry"
	"github.com/nobbmaestro/tmux-tether/pkg/tmux"
	"github.com/spf13/cobra"
)

var sessionAddCmd = &cobra.Command{
	Use:   "add <SESSION> <PATH>",
	Short: "Add tmux session to the storage",
	Args:  cobra.ExactArgs(2),
	RunE:  runSessionAdd,
}

func runSessionAdd(cmd *cobra.Command, args []string) error {
	reg := registry.NewRegistry(registry.WithContext(cmd.Context()))
	ser := reg.GetService()

	s := tmux.Session{
		Name: args[0],
		Path: configpath.Path(args[1]),
	}

	ser.AddToStorage(s)

	return nil
}
