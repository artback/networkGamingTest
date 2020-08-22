package scoreboard

import (
	"github.com/artback/networkGamingTest/internal/jsonwriter"
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScoreboard_AddScore1(t *testing.T) {
	t.Run("Addscore to player", func(t *testing.T) {
		s := Scoreboard{}
		mock := jsonwriter.Mock{Err: nil}
		p := player.NewPlayer("ola", mock, &[2]int{0, 1})
		s.AddResult(5, p)
		require.Equal(t, Scoreboard{"ola": -1}, s)
	})
	t.Run("guess is nil", func(t *testing.T) {
		s := Scoreboard{}
		p := player.NewPlayer("ola", nil, nil)
		err := s.AddResult(5, p)
		require.EqualError(t, err, "guess is nil")
	})
}
