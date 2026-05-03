package cmd

import (
	"bytes"
	"context"
	"fmt"

	"github.com/nobbmaestro/tmux-tether/pkg/config"
	"github.com/nobbmaestro/tmux-tether/pkg/registry"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	printUserConfig    bool
	printDefaultConfig bool
	printConfigPath    bool
)

var rootCmd = &cobra.Command{
	Use:   "tmux-tether [flags]",
	Short: "tmux-tether",
	Args:  cobra.ArbitraryArgs,
	RunE:  runRoot,
}

func runRoot(cmd *cobra.Command, args []string) error {
	switch {
	case printUserConfig:
		return runPrintUserConfig(cmd, args)
	case printDefaultConfig:
		return runPrintDefaultConfig(cmd, args)
	case printConfigPath:
		return runPrintConfigPath(cmd, args)
	default:
		return nil
	}
}

func runPrintUserConfig(cmd *cobra.Command, args []string) error {
	reg := registry.NewRegistry(registry.WithContext(cmd.Context()))
	return printConfig(reg.GetConfig())
}

func runPrintDefaultConfig(cmd *cobra.Command, args []string) error {
	return printConfig(config.GetDefaultUserConfig())
}

func runPrintConfigPath(cmd *cobra.Command, args []string) error {
	reg := registry.NewRegistry(registry.WithContext(cmd.Context()))
	fmt.Println(reg.GetConfigPath())
	return nil
}

func printConfig(cfg *config.UserConfig) error {
	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("error encoding config: %w", err)
	}
	fmt.Printf("%s\n", buf.String())
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
	rootCmd.
		Flags().
		BoolVarP(&printUserConfig, "user-config", "c", false, "print the user config")
	rootCmd.
		Flags().
		BoolVarP(&printDefaultConfig, "default-config", "C", false, "print the default config")
	rootCmd.
		Flags().
		BoolVarP(&printConfigPath, "config-dir", "d", false, "print the config directory")
}
