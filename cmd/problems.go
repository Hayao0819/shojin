package cmd

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/Hayao0819/shojin/problems"
	"github.com/spf13/cobra"
)

func problemsCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "problems",
		Short: "Problems related commands",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			client := problems.NewKenkooooClient()

			switch args[0] {
			case "contest":
				contests, err := client.GetContests()
				if err != nil {
					return err
				}
				for _, contest := range contests {
					cmd.Println(contest.Title)
				}
			case "problem":
				problems, err := client.GetProblems()
				if err != nil {
					return err
				}
				for _, problem := range problems {
					cmd.Println(problem.Title)
				}
			case "problem-detail":
				problems, err := client.GetDetailProblem()
				if err != nil {
					return err
				}
				for _, problem := range problems {
					cmd.Println(problem.Title)
				}
			}

			return nil

		},
	}

	return &cmd
}

func init() {
	cobrautils.AddSubCmds(problemsCmd())
}
