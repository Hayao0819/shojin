package cmd

import (
	"os"
	"path"

	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/Hayao0819/nahi/mobra"
	"github.com/Hayao0819/shojin/conf"
	"github.com/spf13/cobra"
)

func makeCmd() *cobra.Command {
	runPermission := false
	cmd := mobra.New("make ABC001 A").
		PersistentPreRunEWithParent(func(cmd *cobra.Command, args []string) error {
			return conf.Initialize()
		}).
		Args(cobra.RangeArgs(2, 3)).
		RunE(func(cmd *cobra.Command, args []string) error {
			contestDir, err := conf.GetContestDir()
			if err != nil {
				return err
			}

			contest := args[0]
			problem := args[1]

			if err := os.MkdirAll(path.Join(contestDir, contest, problem), 0750); err != nil {
				return err
			}

			file := args[2]
			if file == "" {
				return nil
			}

			filePath := path.Join(contestDir, contest, problem, file)
			fileObj, err := os.Create(filePath)
			if err != nil {
				return err
			}
			if runPermission {
				if err := fileObj.Chmod(0750); err != nil {
					return err
				}
			}

			return nil
		}).Cobra()
	cmd.Flags().BoolVarP(&runPermission, "execute", "x", false, "Add execute permission to the created file")

	return cmd
}

func init() {
	cobrautils.AddSubCmds(makeCmd())
}
