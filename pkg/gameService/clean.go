package gameService

import (
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"log"
)

func (gs *GameService) Clean() {
	gs.players = map[string]player.Player{}
	gs.Scoreboard = scoreboard.Scoreboard{}
	observers := gs.observers
	gs.started = false
	for _, o := range observers {
		gs.join(o)
	}
	log.Print(gs.players)

}
