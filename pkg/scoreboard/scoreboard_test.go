package scoreboard

import (
	"github.com/artback/networkGamingTest/pkg/gamehandler"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScoreboard_AddScore(t *testing.T) {
	t.Run("add scores to an existing player", func(t *testing.T) {
		s := NewScoreBoard(5)
		pOlof := gamehandler.Player{Name: "olof"}
		s.AddScore(2, pOlof)
		s.AddScore(2, pOlof)
		score, _ := s.Get(pOlof)
		require.Equal(t, 4, score)
	})
}

func TestScoreboard_JoinScoreboards(t *testing.T) {
	t.Run("joining two scoreboards together", func(t *testing.T) {
		s1 := NewScoreBoard(1)
		s2 := NewScoreBoard(1)
		pOlof := gamehandler.Player{Name: "olof"}
		s1.AddScore(2, pOlof)
		s2.AddScore(2, pOlof)
		s1.Join(s2)
		score, _ := s1.Get(pOlof)
		require.Equal(t, 4, score)
	})
}
