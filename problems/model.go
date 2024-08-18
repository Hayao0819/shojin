package problems

type Contest struct {
	Id               string `json:"id"`
	StartEpochSecond int    `json:"start_epoch_second"`
	DurationSecond   int    `json:"duration_second"`
	Title            string `json:"title"`
	RateChange       string `json:"rate_change"`
}

type Problem struct {
	Id           string `json:"id"`
	ContestId    string `json:"contest_id"`
	ProblemIndex string `json:"problem_index"`
	Name         string `json:"name"`
	Title        string `json:"title"`
}

type DetailedProblem struct {
	Problem
	ShortestSubmissionId int    `json:"shortest_submission_id"`
	ShortestContestId    string `json:"shortest_contest_id"`
	ShortestUserId       string `json:"shortest_user_id"`
	FastestSubmissionId  int    `json:"fastest_submission_id"`
	FastestContestId     string `json:"fastest_contest_id"`
	FastestUserId        string `json:"fastest_user_id"`
	FirstSubmissionId    int    `json:"first_submission_id"`
	FirstContestId       string `json:"first_contest_id"`
	FirstUserId          string `json:"first_user_id"`
	SourceCodeLength     int    `json:"source_code_length"`
	ExecutionTime        int    `json:"execution_time"`
	//Point                int    `json:"point"`
	SolverCount int `json:"solver_count"`
}
