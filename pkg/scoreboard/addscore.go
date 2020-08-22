package scoreboard

import (
	"errors"
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/artback/networkGamingTest/pkg/scoring"
)

func (s Scoreboard) AddResult(result int, p player.Player) error {
	if p.Guess != nil {
		score := scoring.ScoreGuess(result, *p.Guess)
		if s[p.Name] == nil {
			s[p.Name] = NewResult(p)
		}
		s[p.Name].addScore(score)
		return nil
	}
	return errors.New("guess is nil")
}
