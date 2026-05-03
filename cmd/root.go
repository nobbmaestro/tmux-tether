package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmux-tether [flags]",
	Short: "tmux-tether",
	Args:  cobra.ArbitraryArgs,
	RunE:  runRoot,
}

func runRoot(cmd *cobra.Command, args []string) error {
	fmt.Println("Hello, World!")

	return nil
}

func SetContext(ctx context.Context) {
	rootCmd.SetContext(ctx)
}

func SetVersionInfo(version, commit, date string) {
	rootCmd.Version = version
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
}
