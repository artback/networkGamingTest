package player

import (
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"github.com/artback/networkGamingTest/pkg/scoring"
)

func (p *Player) AddScore(result int, scoreboard *scoreboard.Scoreboard) error {
	score := scoring.ScoreGuess(result, p.Guess)
	scoreboard.AddScore(score, p.Name)
	if p.Jw != nil {
		err := p.Jw.WriteJSON(scoreboard)
		if err != nil {
			return err
		}
	}
	return nil
}
