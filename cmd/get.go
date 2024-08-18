package cmd

import (
	"os"
	"path"
	"strconv"

	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/Hayao0819/shojin/conf"
	"github.com/Hayao0819/shojin/oj"
	"github.com/Hayao0819/shojin/problems"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
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
			contestDir, err := conf.GetContestDir()
			if err != nil {
				return err
			}

			contest := ""
			if len(args) > 0 {
				contest = args[0]
			}
			if contest == "" {
				contests, err := problems.GetContests()
				if err != nil {
					return err
				}

				contestsStr := funk.Map(contests, func(contest problems.Contest) string {
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

				problemsStr := funk.Map(problemlist, func(problem problems.Problem) string {
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

			if err := os.MkdirAll(path.Join(contestDir, contest, problem), 0750); err != nil {
				return err
			}

			examples, err := oj.GetTestCaces(problemObj)
			if err != nil {
				return err
			}
			for i, example := range examples {
				os.WriteFile(path.Join(contestDir, contest, problem, "ex"+strconv.Itoa(i)+".txt"), example, 0640)
			}

			if len(args) > 2 {
				file := args[2]
				filePath := path.Join(contestDir, contest, problem, file)
				fileObj, err := os.Create(filePath)
				if err != nil {
					return err
				}
				if err := fileObj.Chmod(0750); err != nil {
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
