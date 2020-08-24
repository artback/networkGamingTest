package gameService

import (
	"github.com/artback/networkGamingTest/internal/jsonwriter"
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_gameService_startGame(t *testing.T) {
	t.Run("parse duration error", func(t *testing.T) {
		c := config.Configuration{Sleep: "typo", SleepBetween: "0"}
		g := NewGameService()
		_, _, err := g.startGame(c, 0)
		require.EqualError(t, err, ParseDurationError.Error())

	})
	t.Run("Interval is undefined", func(t *testing.T) {
		c := config.Configuration{Sleep: "0s", SleepBetween: "0", Rounds: 1}
		g := NewGameService()
		_, _, err := g.startGame(c, 0)
		require.EqualError(t, err, "interval is undefined")
	})
	t.Run("Run game with 2 players", func(t *testing.T) {
		c := config.Configuration{Sleep: "0s", SleepBetween: "0", Rounds: 1, EndInterval: 10, BeginInterval: 0}
		g := NewGameService()
		mock := jsonwriter.Mock{}
		g.Join("ola", mock, &[2]int{0, 5})
		g.Join("per", mock, &[2]int{2, 3})
		res, winner, err := g.startGame(c, 1)
		require.Nil(t, err)
		ret := *res
		require.Nil(t, winner)
		require.Equal(t, -1, ret["per"].Score)
		require.Equal(t, 0, ret["ola"].Score)
	})
}
