package session

import (
	"github.com/nobbmaestro/tmux-tether/pkg/registry"
	"github.com/nobbmaestro/tmux-tether/pkg/tmux"
	"github.com/spf13/cobra"
)

var sessionRmCmd = &cobra.Command{
	Use:   "rm <SESSION>",
	Short: "Remove tmux session from the storage",
	Args:  cobra.ExactArgs(1),
	RunE:  runSessionRm,
}

func runSessionRm(cmd *cobra.Command, args []string) error {
	reg := registry.NewRegistry(registry.WithContext(cmd.Context()))
	ser := reg.GetService()

	s := tmux.Session{
		Name: args[0],
	}

	ser.RemoveFromStorage(s)

	return nil
}
