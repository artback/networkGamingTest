package gameService

import (
	"github.com/artback/networkGamingTest/internal/jsonwriter"
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_gameService_startGame(t *testing.T) {
	t.Run("parse duration error", func(t *testing.T) {
		c := config.Configuration{Sleep: "typo", SleepBetween: "0"}
		g := NewGameService()
		_, err := g.startGame(c, 0)
		require.EqualError(t, err, ParseDurationError.Error())

	})
	t.Run("Interval is undefined", func(t *testing.T) {
		c := config.Configuration{Sleep: "0s", SleepBetween: "0", Rounds: 1}
		g := NewGameService()
		_, err := g.startGame(c, 0)
		require.EqualError(t, err, "interval is undefined")
	})
	t.Run("Run game with 2 players", func(t *testing.T) {
		c := config.Configuration{Sleep: "0s", SleepBetween: "0", Rounds: 1, EndInterval: 10, BeginInterval: 0}
		g := NewGameService()
		g.Join("ola", nil, &[2]int{0, 5})
		g.Join("per", nil, &[2]int{2, 3})
		res, err := g.startGame(c, 1)
		require.Nil(t, err)
		require.Equal(t, scoreboard.Scoreboard{"ola": 0, "per": -1}, *res)
	})
}

func Test_gameService_Join(t *testing.T) {
	t.Run("join not started game", func(t *testing.T) {
		gs := NewGameService()
		gs.Join("ola", nil, &[2]int{2, 3})
		expected := map[string]player.Player{"ola": player.NewPlayer("ola", jsonwriter.JsonWriter(nil), &[2]int{2, 3})}
		require.Equal(t, expected, gs.players)
	})
	t.Run("join started game", func(t *testing.T) {
		gs := NewGameService()
		gs.started = true
		gs.Join("ola", nil, &[2]int{2, 3})
		expected := map[string]player.Player{"ola": player.NewPlayer("ola", jsonwriter.JsonWriter(nil), &[2]int{2, 3})}
		require.Equal(t, expected, gs.observers)
	})
	t.Run("join as observer", func(t *testing.T) {
		gs := NewGameService()
		gs.started = false
		gs.Join("ola", nil, nil)
		expected := map[string]player.Player{"ola": player.NewPlayer("ola", jsonwriter.JsonWriter(nil), nil)}
		require.Equal(t, expected, gs.observers)
	})
}
