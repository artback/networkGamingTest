package gameService

import (
	"errors"
	"github.com/artback/networkGamingTest/internal/jsonwriter"
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/artback/networkGamingTest/pkg/game"
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"time"
)

var ParseDurationError = errors.New("couldn't parse duration string")

type gameService struct {
	players   []player.Player
	observers []player.Player
	scoreboard.Scoreboard
	started bool
}

func NewGameService() *gameService {
	return &gameService{}
}

// Runs the game in a loop that sleeps for time set by SleepBetween from Configuration
// cleans the players after game and and move the observers as players for the next game
func (gs *gameService) GameLoop(c config.GameConfiguration, seed int64, total *scoreboard.Scoreboard) (*scoreboard.Scoreboard, error) {
	sleep, err := time.ParseDuration(c.SleepBetween)
	if err != nil {
		return nil, ParseDurationError
	}
	for {
		time.Sleep(sleep)
		if c.MinimumPlayer >= len(gs.players) {
			board, err := gs.startGame(c, seed)
			if err != nil {
				return board, err
			}
			total.Join(board)
			for _, p := range gs.players {
				err = p.Jw.Close()
				if err != nil {
					return nil, err
				}
			}
			gs.resetGameService()

		}
	}
}
func (gs *gameService) resetGameService() {
	observers := gs.observers
	g := new(gameService)
	for _, o := range observers {
		g.join(o)
	}
}

// Runs one game , keeps track of score and sends updates to players
func (gs *gameService) startGame(c config.GameConfiguration, seed int64) (*scoreboard.Scoreboard, error) {
	gs.started = true
	defer func() {
		gs.started = false
	}()

	board := scoreboard.NewScoreBoard(len(gs.players))
	sleep, err := time.ParseDuration(c.Sleep)
	if err != nil {
		return board, ParseDurationError
	}
	g := game.NewGame(c.Rounds, seed, sleep, [2]int{c.BeginInterval, c.EndInterval})
	resChan := make(chan int)
	errChan := make(chan error)
	go g.Run(resChan, errChan)
	for {
		select {
		case result, ok := <-resChan:
			if !ok {
				return board, nil
			}
			gs.addToScoreBoard(result, board)
		case err := <-errChan:
			return board, err
		}
	}
}
func (gs *gameService) addToScoreBoard(result int, scoreboard *scoreboard.Scoreboard) {
	for _, p := range append(gs.players, gs.observers...) {
		p.AddScore(result, scoreboard)
	}
}

// Method for joining game. Will place user in players if no game is currently active.
// If no game is active user will join player and if guess is zero
//or game already started user joins observers
func (gs *gameService) Join(name string, ws jsonwriter.JsonWriter, guess [2]int) {
	p := player.NewPlayer(name, ws, guess)
	gs.join(p)
}
func (gs *gameService) join(p player.Player) {
	if !gs.started && p.Guess[0] >= 0 && p.Guess[1] > 0 {
		gs.players = append(gs.players, p)
	} else {
		gs.observers = append(gs.observers, p)
	}
}
