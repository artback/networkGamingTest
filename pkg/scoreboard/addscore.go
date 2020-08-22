package scoreboard

import (
	"errors"
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/artback/networkGamingTest/pkg/scoring"
)

func (s Scoreboard) AddResult(result int, p player.Player) error {
	if p.Guess != nil {
		score := scoring.ScoreGuess(result, *p.Guess)
		s[p.Name] = s[p.Name] + score
		return nil
	}
	return errors.New("guess is nil")

}
