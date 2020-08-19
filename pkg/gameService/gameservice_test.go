package gameService

import (
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_gameService_startGame(t *testing.T) {
	t.Run("parse duration error", func(t *testing.T) {
		c := config.GameConfiguration{Sleep: "typo", SleepBetween: "0"}
		g := NewGameService()
		_, err := g.startGame(c, 0)
		require.EqualError(t, err, ParseDurationError.Error())

	})
	t.Run("Interval is undefined", func(t *testing.T) {
		c := config.GameConfiguration{Sleep: "0s", SleepBetween: "0", Rounds: 1}
		g := NewGameService()
		_, err := g.startGame(c, 0)
		require.EqualError(t, err, "interval is undefined")
	})
	t.Run("Run game with 2 players", func(t *testing.T) {
		c := config.NewConfig()
		g := NewGameService()
		g.Join("ola", nil, [2]int{0, 5})
		g.Join("per", nil, [2]int{2, 3})
		res, _ := g.startGame(c.Game, 1)
		require.Equal(t, map[string]int{"ola": 0, "per": -1}, res.Map())

	})
}

func Test_gameService_Join(t *testing.T) {
	t.Run("join not started game", func(t *testing.T) {
		gs := NewGameService()
		gs.Join("ola", nil, [2]int{2, 3})
		expected := []player.Player{player.NewPlayer("ola", nil, [2]int{2, 3})}
		require.Equal(t, expected, gs.players)
	})
	t.Run("join started game", func(t *testing.T) {
		gs := NewGameService()
		gs.started = true
		gs.Join("ola", nil, [2]int{2, 3})
		expected := []player.Player{player.NewPlayer("ola", nil, [2]int{2, 3})}
		require.Equal(t, expected, gs.observers)
	})
	t.Run("join as observer", func(t *testing.T) {
		gs := NewGameService()
		gs.started = false
		gs.Join("ola", nil, [2]int{0, 0})
		expected := []player.Player{player.NewPlayer("ola", nil, [2]int{0, 0})}
		require.Equal(t, expected, gs.observers)
	})
}

func Test_gameService_resetGameService(t *testing.T) {
	t.Run("reset game and close players", func(t *testing.T) {
		gs := NewGameService()
		gs.Join("ola", nil, [2]int{0, 0})
		gs.resetGameService()
		require.Equal(t, []player.Player(nil), gs.players)
	})
}
