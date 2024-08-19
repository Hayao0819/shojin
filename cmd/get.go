package cmd

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/Hayao0819/shojin/code"
	"github.com/Hayao0819/shojin/conf"
	"github.com/Hayao0819/shojin/problems"
	"github.com/manifoldco/promptui"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func getCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "get",
		Short: "Get problem",
		Args:  cobra.MaximumNArgs(3),
		PersistentPreRunE: cobrautils.WithParentPersistentPreRunE(func(cmd *cobra.Command, args []string) error {
			return conf.Initialize()
		}),
		RunE: func(cmd *cobra.Command, args []string) error {
			contest := ""
			if len(args) > 0 {
				contest = args[0]
			}
			if contest == "" {
				contests, err := problems.GetContests()
				if err != nil {
					return err
				}

				contestsStr := lo.Map(contests, func(contest problems.Contest, index int) string {
					return contest.Title
				})

				prompt := promptui.Select{
					Label: "Select contest",
					Items: contestsStr,
				}

				_, result, err := prompt.Run()
				if err != nil {
					return err
				}
				contest = result
			}

			problem := ""
			if len(args) > 1 {
				problem = args[1]
			}
			if problem == "" {
				problemlist, err := problems.GetContestProblems(contest)
				if err != nil {
					return err
				}

				problemsStr := lo.Map(problemlist, func(problem problems.Problem, index int) string {
					return problem.Title
				})

				prompt := promptui.Select{
					Label: "Select problem",
					Items: problemsStr,
				}

				_, result, err := prompt.Run()
				if err != nil {
					return err
				}
				problem = result
			}

			problemObj, err := problems.GetProblem(contest, problem)
			if err != nil {
				return err
			}

			if _, err := code.CreateDir(problemObj); err != nil {
				return err
			}

			if _, err := code.FetchTestCases(problemObj); err != nil {
				return err
			}

			if len(args) > 2 {
				file := args[2]
				if _, err := code.CreateCode(problemObj, file); err != nil {
					return err
				}
			}

			return nil

		},
	}
	return &cmd
}

func init() {
	cobrautils.AddSubCmds(getCmd())
}
