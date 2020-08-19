package scoreboard_test

import (
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScoreboard_AddScore(t *testing.T) {
	t.Run("add scores to an existing player", func(t *testing.T) {
		s := scoreboard.NewScoreBoard(5)
		pOlof := player.Player{Name: "olof"}
		s.AddScore(2, pOlof.Name)
		s.AddScore(2, pOlof.Name)
		score, _ := s.Get(pOlof.Name)
		require.Equal(t, 4, score)
	})
}

func TestScoreboard_JoinScoreboards(t *testing.T) {
	t.Run("joining two scoreboards together", func(t *testing.T) {
		s1 := scoreboard.NewScoreBoard(1)
		s2 := scoreboard.NewScoreBoard(1)
		pOlof := player.Player{Name: "olof"}
		s1.AddScore(2, pOlof.Name)
		s2.AddScore(2, pOlof.Name)
		s1.Join(s2)
		score, _ := s1.Get(pOlof.Name)
		require.Equal(t, 4, score)
	})
}

func TestScoreboard_MapScoreboards(t *testing.T) {
	t.Run("return a map from a scoreboard", func(t *testing.T) {
		pOlof := player.Player{Name: "olof"}
		s1 := scoreboard.NewScoreBoard(1)
		s1.AddScore(2, pOlof.Name)
		s1.AddScore(2, pOlof.Name)
		scores := s1.Map()
		require.Equal(t, map[string]int{"olof": 4}, scores)
	})
}
