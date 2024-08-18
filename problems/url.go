package problems

func (p *Problem) GetUrl() string {
	return "https://atcoder.jp/contests/" + p.ContestId + "/tasks/" + p.Id
}
