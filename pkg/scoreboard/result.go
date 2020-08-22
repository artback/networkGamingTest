package scoreboard

import "github.com/artback/networkGamingTest/pkg/player"

type Result struct {
	Name  string `json:"name"`
	guess *[2]int
	Score int `json:"score"`
}

func (r *Result) addScore(score int) {
	if r != nil {
		r.Score = r.Score + score
	}
}

func NewResult(p player.Player) *Result {
	return &Result{Name: p.Name, guess: p.Guess}
}
