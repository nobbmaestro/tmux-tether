package switcher

import (
	"errors"
	"strconv"

	"github.com/nobbmaestro/tmux-tether/pkg/registry"
	"github.com/nobbmaestro/tmux-tether/pkg/shell"
	"github.com/nobbmaestro/tmux-tether/pkg/tmux"
	"github.com/spf13/cobra"
)

var SwitchCmd = &cobra.Command{
	Use:   "switch [MARK]",
	Short: "Switch to a marked tmux session",
	Args:  cobra.ArbitraryArgs,
	RunE:  runSwitch,
}

var ErrInvalidID = errors.New("no session at provided mark")

func runSwitch(cmd *cobra.Command, args []string) error {
	reg := registry.NewRegistry(registry.WithContext(cmd.Context()))
	cfg := reg.GetConfig()

	t := tmux.New(
		shell.New(),
		cfg.TmuxCommand,
	)

	if len(args) == 0 {
		return ErrInvalidID
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	if id >= len(cfg.Switch.Markers) {
		return ErrInvalidID
	}

	err = t.CreateOrSwitch(cfg.Switch.Markers[id])
	if err != nil {
		return err
	}

	return nil
}
