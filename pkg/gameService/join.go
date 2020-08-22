package gameService

import (
	"github.com/artback/networkGamingTest/internal/jsonwriter"
	"github.com/artback/networkGamingTest/pkg/player"
)

// Method for users to join game.
// If no game is active user will join players
//if guess is zero or game already started user joins observers
func (gs *GameService) Join(name string, ws jsonwriter.JsonWriter, guess *[2]int) {
	p := player.NewPlayer(name, ws, guess)
	gs.join(p)
}
func (gs *GameService) join(p player.Player) {
	if !gs.started && p.Guess != nil {
		gs.players[p.Name] = p
	} else {
		gs.observers[p.Name] = p
	}
}
