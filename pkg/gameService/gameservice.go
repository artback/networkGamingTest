package gameService

import (
	"errors"
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/artback/networkGamingTest/pkg/game"
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"time"
)

var ParseDurationError = errors.New("couldn't parse duration string")

type GameService struct {
	players   map[string]player.Player
	observers map[string]player.Player
	scoreboard.Scoreboard
	total   scoreboard.Scoreboard
	started bool
}

func NewGameService() *GameService {
	return &GameService{
		players:   map[string]player.Player{},
		observers: map[string]player.Player{},
	}
}

// Runs one game , keeps track of score and sends updates to players
func (gs *GameService) startGame(c config.Configuration, seed int64) (*scoreboard.Scoreboard, *scoreboard.Result, error) {
	gs.started = true
	defer func() {
		gs.started = false
	}()

	board := &scoreboard.Scoreboard{}
	sleep, err := time.ParseDuration(c.Sleep)
	if err != nil {
		return board, nil, ParseDurationError
	}
	g := game.NewGame(c.Rounds, seed, sleep, [2]int{c.BeginInterval, c.EndInterval})
	resChan := make(chan int)
	errChan := make(chan error)
	go g.Run(resChan, errChan)
	for {
		select {
		case result, ok := <-resChan:
			if !ok {
				return board, nil, nil
			}
			err := gs.addToScoreBoard(result, board)
			if res := board.Have21(); res != nil {
				return board, res, nil
			}
			if err != nil {
				return board, nil, err
			}
		case err := <-errChan:
			return board, nil, err
		}
	}
}
func (gs *GameService) addToScoreBoard(result int, scoreboard *scoreboard.Scoreboard) error {
	if scoreboard != nil {
		for _, p := range gs.players {
			if err := scoreboard.AddResult(result, p); err != nil {
				return err
			}
			if err := p.WriteMessage("result", scoreboard); err != nil {
				return err
			}
		}
		for _, o := range gs.observers {
			if err := o.WriteMessage("result", scoreboard); err != nil {
				return err
			}
		}
	}
	return nil
}
