package scoreboard

import (
	"sort"
)

func (s *Scoreboard) Winner() *Result {
	var results []Result
	for _, result := range *s {
		results = append(results, *result)
	}
	sort.Slice(results, func(i, j int) bool {
		rI := results[i]
		rJ := results[j]
		if rI.Score != rJ.Score {
			return rI.Score > rJ.Score
		} else if rI.guess != nil && rJ.guess != nil {
			if rI.guess[1] != rJ.guess[1] {
				return rI.guess[1] > rJ.guess[1]
			} else if rI.guess[0] != rJ.guess[0] {
				return rI.guess[0] > rJ.guess[0]
			}
		} else if rI.Name != rJ.Name {
			return rI.Name < rJ.Name
		}
		return false
	})
	return &results[0]
}
