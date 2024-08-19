package cmd

import (
	"errors"
	"strings"

	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/Hayao0819/shojin/code"
	"github.com/Hayao0819/shojin/conf"
	"github.com/Hayao0819/shojin/problems"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func randomCmd() *cobra.Command {
	contest := ""

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

			if contest != "" {
				contest = strings.ToLower(contest)
				println(contest)
				filtered = lo.Filter(filtered, func(p problems.Problem, index int) bool {
					pLowerContest := strings.ToLower(p.ContestId)
					println(pLowerContest)
					return strings.HasPrefix(pLowerContest, contest)
				})
			}

			if len(filtered) == 0 {
				return errors.New("no problems found")
			}

			selected := lo.Sample(filtered)

			cmd.Println(selected.Title)

			if _, err := code.CreateDir(&selected); err != nil {
				return err
			}

			if _, err := code.FetchTestCases(&selected); err != nil {
				return err
			}

			cmd.Println("Problem created successfully")
			cmd.Println("Please open the following URL in your browser:")
			cmd.Println(selected.GetUrl())

			return nil
		},
	}
	cmd.Flags().StringVarP(&contest, "contest", "c", "", "contest prefix")

	return &cmd
}

func init() {
	cobrautils.AddSubCmds(randomCmd())
}
