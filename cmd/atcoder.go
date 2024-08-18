package cmd

import (
	"github.com/Hayao0819/nahi/cobrautils"
	"github.com/Hayao0819/shojin/oj"
	"github.com/Hayao0819/shojin/problems"
	"github.com/spf13/cobra"
)

var atcoderSubCmds = cobrautils.Registory{}

func testcaseCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:  "testcase",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			targetContest := args[0]
			targetProblem := args[1]

			problem, err := problems.GetProblem(targetContest, targetProblem)
			if err != nil {
				return err
			}

			testcases, err := oj.GetTestCaces(problem)
			if err != nil {
				return err
			}

			for _, testcase := range testcases {
				cmd.Println("---")
				cmd.Print(string(testcase))
			}

			return nil
		},
	}
	return &cmd
}

func atcoderCmd() *cobra.Command {
	cmd := cobra.Command{
		Use: "atcoder contest problem",
	}
	atcoderSubCmds.Bind(&cmd)
	return &cmd
}

func init() {
	atcoderSubCmds.Add(testcaseCmd())
	cobrautils.AddSubCmds(atcoderCmd())
}
