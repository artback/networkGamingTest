package gameService

import (
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGameService_Clean(t *testing.T) {
	t.Run("clean with two players", func(t *testing.T) {
		gs := NewGameService()
		gs.Join("ola", nil, &[2]int{2, 3})
		gs.Join("bjorn", nil, &[2]int{2, 3})
		gs.Clean()
		require.Equal(t, map[string]player.Player{}, gs.players)
		require.Equal(t, map[string]player.Player{}, gs.observers)
		require.Equal(t, scoreboard.Scoreboard{}, gs.Scoreboard)
	})
	t.Run("clean with two players", func(t *testing.T) {
		gs := NewGameService()
		gs.started = true
		p1 := player.NewPlayer("ola", nil, &[2]int{2, 3})
		p2 := player.NewPlayer("bjorn", nil, &[2]int{2, 3})
		gs.join(p1)
		gs.join(p2)
		gs.Clean()
		require.Equal(t, map[string]player.Player{"ola": p1, "bjorn": p2}, gs.players)
		require.Equal(t, map[string]player.Player{}, gs.observers)
		require.Equal(t, scoreboard.Scoreboard{}, gs.Scoreboard)
	})
}
