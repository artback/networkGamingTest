package gameService

import (
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"log"
	"time"
)

// Runs the game in a loop that sleeps for time set by SleepBetween from Configuration
// cleans the players after game and and move the observers as players for the next game
func GameLoop(gs *GameService, c config.Configuration, seed int64) {
	sleep, err := time.ParseDuration(c.SleepBetween)
	if err != nil {
		return
	}
	for {
		time.Sleep(sleep)
		if c.MinimumPlayer <= len(gs.players) {
			board, err := gs.startGame(c, seed)
			if err != nil {
				log.Print(err)
			}
			gs.total = scoreboard.Join(gs.total, *board)
			for _, p := range gs.players {
				err := p.WriteMessage("board", board)
				if err != nil {
					log.Print(err)
				}
				err = p.WriteMessage("total", gs.total)
				if err != nil {
					log.Print(err)
				}
				err = p.WriteMessage("winner", board.Winner())
				if err != nil {
					log.Print(err)
				}
			}
			for _, o := range gs.observers {
				err := o.WriteMessage("board", board)
				if err != nil {
					log.Print(err)
				}
				err = o.WriteMessage("total", gs.total)
				if err != nil {
					log.Print(err)
				}
				err = o.WriteMessage("winner", board.Winner())
				if err != nil {
					log.Print(err)
				}
			}
			gs.Clean()
		}
		gs.PingPlayers()
	}
}
