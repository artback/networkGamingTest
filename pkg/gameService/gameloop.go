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
			log.Print("New game started with ", len(gs.players), " players")
			board, winner, err := gs.startGame(c, seed)
			if err != nil {
				log.Print("StartGame", err)
			}
			log.Print("Game ended without errors, winner: ", board.Winner().Name)
			gs.total = scoreboard.Join(gs.total, *board)
			if winner == nil {
				winner = board.Winner()
			}
			gs.announceResult(board, winner)
			gs.Clean()
		}
		gs.PingPlayers()
	}
}

func (gs *GameService) announceResult(board *scoreboard.Scoreboard, winner *scoreboard.Result) {
	for _, p := range gs.players {
		err := p.WriteMessage("result", board)
		if err != nil {
			log.Print("Error writing to ", p.Name, " ", err)
		}
		err = p.WriteMessage("total", gs.total)
		if err != nil {
			log.Print("Error writing to ", p.Name, " ", err)
		}
		err = p.WriteMessage("winner", winner)
		if err != nil {
			log.Print("Error writing to ", p.Name, " ", err)
		}
	}
	for _, o := range gs.observers {
		err := o.WriteMessage("result", board)
		if err != nil {
			log.Print("Error writing to ", o.Name, " ", err)
		}
		err = o.WriteMessage("total", gs.total)
		if err != nil {
			log.Print("Error writing to ", o.Name, " ", err)
		}
		err = o.WriteMessage("winner", winner)
		if err != nil {
			log.Print("Error writing to ", o.Name, " ", err)
		}
	}

}
