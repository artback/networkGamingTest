package gameService

import (
	"github.com/artback/networkGamingTest/internal/jsonwriter"
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/stretchr/testify/require"
	"testing"
)

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
