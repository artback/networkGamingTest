package gameService

import (
	"github.com/artback/networkGamingTest/pkg/player"
	"testing"
)

func TestGameService_PingPlayers(t *testing.T) {
	t.Run("nil players and observers", func(t *testing.T) {
		gs := &GameService{}
		gs.PingPlayers()
	})
	t.Run("nil players and observers", func(t *testing.T) {
		gs := NewGameService()
		gs.join(player.NewPlayer("ola", nil, nil))
		gs.PingPlayers()
	})
}
