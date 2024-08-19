package cmd

import (
	"strings"

	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/Hayao0819/shojin/code"
	"github.com/Hayao0819/shojin/conf"
	"github.com/Hayao0819/shojin/problems"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func randomCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "random difficulty",
		Args: cobra.ExactArgs(1),
		PersistentPreRunE: cobrautils.WithParentPersistentPreRunE(func(cmd *cobra.Command, args []string) error {
			return conf.Initialize()
		}),
		RunE: func(cmd *cobra.Command, args []string) error {
			difficulty := args[0]

			list, err := problems.GetProblems()
			if err != nil {
				return err
			}

			filtered := lo.Filter(list, func(p problems.Problem, index int) bool {
				return strings.EqualFold(p.ProblemIndex, difficulty)
			})

			selected := lo.Sample(filtered)
			cmd.Println(selected.Title)

			if _, err := code.CreateDir(&selected); err != nil {
				return err
			}

			if _, err := code.FetchTestCases(&selected); err != nil {
				return err
			}

			return nil
		},
	}

	return &cmd
}

func init() {
	cobrautils.AddSubCmds(randomCmd())
}
