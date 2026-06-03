package session

import "github.com/spf13/cobra"

var SessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Manage tracked tmux sessions",
}

func init() {
	SessionCmd.AddCommand(sessionAddCmd)
	SessionCmd.AddCommand(sessionRmCmd)
}
