package clone

import (
	"cmp"
	"os"
	"path"
	"strings"

	configpath "github.com/nobbmaestro/tmux-tether/pkg/config/parsers"
	"github.com/nobbmaestro/tmux-tether/pkg/git"
	"github.com/nobbmaestro/tmux-tether/pkg/registry"
	"github.com/nobbmaestro/tmux-tether/pkg/tmux"
	"github.com/spf13/cobra"
)

type CloneOptions struct {
	dir  string
	name string
}

var cloneOpts = &CloneOptions{}

var CloneCmd = &cobra.Command{
	Use:   "clone [flags] <URL>",
	Short: "Clone repo from URL and initialize session",
	Args:  cobra.ExactArgs(1),
	RunE:  runClone,
}

func nameFromGitUrl(url string) string {
	name := path.Base(url)
	return strings.TrimSuffix(name, ".git")
}

func runClone(cmd *cobra.Command, args []string) error {
	reg := registry.NewRegistry(registry.WithContext(cmd.Context()))
	cfg := reg.GetConfig()
	ser := reg.GetService()

	repoName := nameFromGitUrl(args[0])

	parentDir := cmp.Or(
		cloneOpts.dir,
		cfg.Session.Dirs[0].String())

	sessionName := cmp.Or(
		cloneOpts.name,
		repoName)

	g := git.New(
		git.WithBin(cfg.GitCommand),
	)

	err := os.MkdirAll(parentDir, 0755)
	if err != nil {
		return err
	}

	err = os.Chdir(parentDir)
	if err != nil {
		return err
	}

	err = g.Clone(args[0])
	if err != nil {
		return err
	}

	err = ser.CreateOrSwitch(tmux.Session{
		Name: sessionName,
		Path: configpath.New(path.Join(parentDir, repoName)),
	})
	if err != nil {
		return err
	}

	return nil
}

func init() {
	CloneCmd.
		Flags().
		StringVarP(&cloneOpts.dir, "dir", "d", "", "directory to clone the repository int")
	CloneCmd.
		Flags().
		StringVarP(&cloneOpts.name, "name", "n", "", "session name")
}
