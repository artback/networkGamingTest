package scoreboard

type Result struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func (s *Scoreboard) Winner() *Result {
	var winner Result
	for player, score := range *s {
		if score > winner.Score {
			winner = Result{Name: player, Score: score}
		}
	}
	if winner == (Result{}) {
		return nil
	}
	return &winner
}
