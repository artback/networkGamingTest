package gameService

import (
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
)

func (gs *GameService) Clean() {
	observers := gs.observers
	gs.players = map[string]player.Player{}
	gs.observers = map[string]player.Player{}

	gs.Scoreboard = scoreboard.Scoreboard{}
	gs.started = false
	for _, o := range observers {
		gs.join(o)
	}
}
