package scoreboard

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScoreboard_Winner(t *testing.T) {
	t.Run("get winner", func(t *testing.T) {
		s := Scoreboard{"ola": 5, "per": 1, "sven": 3}
		winner := s.Winner()
		require.Equal(t, &Result{Name: "ola", Score: 5}, winner)
	})
	t.Run("nil scoreboard", func(t *testing.T) {
		var s Scoreboard
		winner := s.Winner()
		var expected *Result
		require.Equal(t, expected, winner)
	})
}
