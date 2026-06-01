package last

import (
	"github.com/nobbmaestro/tmux-tether/pkg/registry"
	"github.com/nobbmaestro/tmux-tether/pkg/shell"
	"github.com/nobbmaestro/tmux-tether/pkg/tmux"
	"github.com/spf13/cobra"
)

var LastCmd = &cobra.Command{
	Use:   "last",
	Short: "Switch to last tmux session",
	Args:  cobra.ArbitraryArgs,
	RunE:  runLast,
}

func runLast(cmd *cobra.Command, args []string) error {
	reg := registry.NewRegistry(registry.WithContext(cmd.Context()))
	cfg := reg.GetConfig()

	t := tmux.New(
		shell.New(),
		cfg.TmuxCommand,
	)

	err := t.SwitchLastSession()
	if err != nil {
		return err
	}

	return nil
}
