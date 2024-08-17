package cmd

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/Hayao0819/shojin/code"
	"github.com/Hayao0819/shojin/conf"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "list",
		PersistentPreRunE: cobrautils.WithParentPersistentPreRunE(func(cmd *cobra.Command, args []string) error {
			return conf.Initialize()
		}),
		RunE: func(cmd *cobra.Command, args []string) error {
			contests, err := code.GetContestList()
			if err != nil {
				return err
			}
			for _, c := range *contests {
				cmd.Println(c)
			}
			return nil
		},
	}
	return &cmd
}

func init() {
	cobrautils.AddSubCmds(listCmd())
}
