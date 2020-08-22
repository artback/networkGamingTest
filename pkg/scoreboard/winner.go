package scoreboard

func (s *Scoreboard) Winner() *Result {
	var winner *Result
	for _, result := range *s {
		if winner == nil || result.Score > winner.Score {
			winner = result
		}
	}
	return winner
}
