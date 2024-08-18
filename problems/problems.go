package problems

import "strings"

var problemsCache []Problem = nil

func (c *Client) GetProblems() ([]Problem, error) {
	var problems []Problem
	if problemsCache != nil {
		return problemsCache, nil
	}
	if err := c.FetchAndUnmarshal("problems", &problems); err != nil {
		return nil, err
	}
	problemsCache = problems
	return problems, nil
}

func (c *Client) GetContestProblems(contestId string) ([]Problem, error) {
	problems, err := client.GetProblems()
	if err != nil {
		return nil, err
	}

	var filteredProblems []Problem
	for _, problem := range problems {
		if strings.EqualFold(problem.ContestId, contestId) {
			filteredProblems = append(filteredProblems, problem)
		}
	}
	return filteredProblems, nil
}

func (c *Client) GetProblem(contestId, index string) (*Problem, error) {
	problems, err := c.GetContestProblems(contestId)
	if err != nil {
		return nil, err
	}
	for _, problem := range problems {
		if strings.EqualFold(problem.ProblemIndex, index) {
			return &problem, nil
		}
	}
	return nil, nil
}

func (c *Client) GetDetailProblem() ([]DetailedProblem, error) {
	var problem []DetailedProblem
	if err := c.FetchAndUnmarshal("merged-problems", &problem); err != nil {
		return nil, err
	}
	return problem, nil
}
