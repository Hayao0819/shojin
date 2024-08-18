package problems

import "strings"

var contestsCache []Contest = nil

func (c *Client) GetContests() ([]Contest, error) {
	var contests []Contest
	if contestsCache != nil {
		return contestsCache, nil
	}
	if err := c.FetchAndUnmarshal("contests", &contests); err != nil {
		return nil, err
	}
	contestsCache = contests
	return contests, nil
}

func (c *Client) GetContest(name string) (*Contest, error) {
	contests, err := c.GetContests()
	if err != nil {
		return nil, err
	}
	for _, contest := range contests {
		if strings.EqualFold(contest.Id, name) {
			return &contest, nil
		}
	}
	return nil, nil
}
