package cmd

import (
	"os"
	"path"

	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/Hayao0819/nahi/mobra"
	"github.com/Hayao0819/shojin/conf"
	"github.com/spf13/cobra"
)

func newCmd() *cobra.Command {
	cmd := mobra.New("new ABC001 A").
		PersistentPreRunEWithParent(func(cmd *cobra.Command, args []string) error {
			return conf.Initialize()
		}).
		Args(cobra.ExactArgs(2)).
		RunE(func(cmd *cobra.Command, args []string) error {
			contestDir, err := conf.GetContestDir()
			if err != nil {
				return err
			}

			contest := args[0]
			problem := args[1]

			return os.MkdirAll(path.Join(contestDir, contest, problem), 0750)
		})

	return cmd.Cobra()
}

func init() {
	cobrautils.AddSubCmds(newCmd())
}
