package code

import (
	"os"
	"path"
	"strconv"

	"github.com/Hayao0819/shojin/conf"
	"github.com/Hayao0819/shojin/oj"
	"github.com/Hayao0819/shojin/problems"
)

func ProblemDir(problem *problems.Problem) (string, error) {
	contestDir, err := conf.GetContestDir()
	if err != nil {
		return "", err
	}
	return path.Join(contestDir, problem.ContestId, problem.ProblemIndex), nil
}

func CreateDir(problem *problems.Problem) (string, error) {
	problemDir, err := ProblemDir(problem)
	if err != nil {
		return "", err
	}

	if err := os.MkdirAll(problemDir, 0750); err != nil {
		return "", err
	}

	return problemDir, nil
}

func FetchTestCases(problem *problems.Problem) ([]string, error) {
	problemDir, err := ProblemDir(problem)
	if err != nil {
		return nil, err
	}

	examples, err := oj.GetTestCaces(problem)
	if err != nil {
		return nil, err
	}

	created := []string{}
	for i, example := range examples {
		f := path.Join(problemDir, "ex"+strconv.Itoa(i)+".txt")

		if err := os.WriteFile(f, example, 0640); err == nil {
			created = append(created, f)
		}
	}
	return created, nil
}

func CreateCode(problem *problems.Problem, code string) (string, error) {
	problemDir, err := ProblemDir(problem)
	if err != nil {
		return "", err
	}

	filePath := path.Join(problemDir, code)
	fileObj, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	if err := fileObj.Chmod(0750); err != nil {
		return "", err
	}
	return filePath, nil
}
