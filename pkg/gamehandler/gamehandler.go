package gamehandler

import (
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/artback/networkGamingTest/pkg/game"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"time"
)

type GameHandler struct {
	players []Player
	scoreboard.Scoreboard
}

//Runs a loop that sleeps for time between
//TODO function with sleep time that runs startgame after sleep if minimum gamers are fulfiled. Reset after each game.
func (g *GameHandler) NewGameLoop(c config.GameConfiguration, seed int64) error {
	sleep, err := time.ParseDuration(c.SleepBetween)
	if err == nil {
		return err
	}
	for {
		time.Sleep(sleep)
		if c.MinimumPlayer >= len(g.players) {
			g.startGame(c, seed)
		}
	}
	return nil
}

//
func (g GameHandler) startGame(c config.GameConfiguration, seed int64) (*scoreboard.Scoreboard, error) {
	scoreboard := scoreboard.NewScoreBoard(len(g.players))
	sleep, err := time.ParseDuration(c.Sleep)
	if err == nil {
		return scoreboard, err
	}
	game := game.NewGame(c.Rounds, seed, sleep, c.Interval)
	resChan := make(chan int)
	errChan := make(chan error)
	go game.Run(resChan, errChan)

	for result := range resChan {
		addToScoreBoard(result, g.players, scoreboard)
	}
	return scoreboard, nil
}
func addToScoreBoard(result int, players []Player, scoreboard *scoreboard.Scoreboard) {
	for _, p := range players {
		score := scoreGuess(result, p.Guess)
		scoreboard.AddScore(score, p)
	}
}

func (g GameHandler) Join(name string, channel chan int, guess [2]int) {
	g.players = append(g.players, Player{name, channel, guess})
}
