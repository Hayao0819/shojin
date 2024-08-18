package problems

var client *Client = NewKenkooooClient()

func GetProblems() ([]Problem, error) {
	return client.GetProblems()
}

func GetContestProblems(contestId string) ([]Problem, error) {
	return client.GetContestProblems(contestId)
}

func GetProblem(contestId, index string) (*Problem, error) {
	return client.GetProblem(contestId, index)
}

func GetDetailProblem() ([]DetailedProblem, error) {
	return client.GetDetailProblem()
}

func GetContests() ([]Contest, error) {
	return client.GetContests()
}

func GetContest(name string) (*Contest, error) {
	return client.GetContest(name)
}
