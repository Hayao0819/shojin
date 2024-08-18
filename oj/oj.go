package oj

import (
	"os"
	"os/exec"

	"github.com/Hayao0819/nahi/exutils"
	"github.com/Hayao0819/shojin/flist"
	"github.com/Hayao0819/shojin/problems"
)

func Check() error {
	if exutils.CommandExists("oj") {
		return nil
	}
	return exec.ErrNotFound
}

func GetTestCaces(problem *problems.Problem) ([][]byte, error) {
	tmpdir, err := os.MkdirTemp("", "shojin")
	if err != nil {
		return nil, err
	}
	cmd := exutils.CommandWithStdio("oj", "d", problem.GetUrl(), "-d", tmpdir)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	inputFiles, err := flist.Get(tmpdir, flist.WithExtOnly(".in"))
	if err != nil {
		return nil, err
	}

	var rt [][]byte
	for _, f := range *inputFiles {
		b, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}
		rt = append(rt, b)
	}
	return rt, nil
}
