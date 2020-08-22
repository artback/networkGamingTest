package scoreboard_test

import (
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScoreboard_JoinScoreboards(t *testing.T) {
	t.Run("joining two scoreboards together", func(t *testing.T) {
		s1 := scoreboard.Scoreboard{}
		s2 := scoreboard.Scoreboard{}
		pOlof := player.Player{Name: "olof", Guess: &[2]int{0, 1}}
		err := s1.AddResult(2, pOlof)
		require.Nil(t, err, "error is not nil")
		err = s2.AddResult(2, pOlof)
		require.Nil(t, err, "error is not nil")
		sJoin := scoreboard.Join(s1, s2)
		olof := sJoin[pOlof.Name]
		require.Equal(t, -2, olof.Score)
	})
	t.Run("joining nil scoreboard with non nil", func(t *testing.T) {
		var s1 scoreboard.Scoreboard
		s2 := scoreboard.Scoreboard{"ola": &scoreboard.Result{Name: "ola", Score: 5}}
		pOlof := player.Player{Name: "olof", Guess: &[2]int{0, 1}}
		err := s2.AddResult(2, pOlof)
		require.Nil(t, err, "error is not nil")
		s1 = scoreboard.Join(nil, s2)
		olof := s1[pOlof.Name]
		require.Equal(t, -1, olof.Score)
	})
	t.Run("joining two nil scoreboard ", func(t *testing.T) {
		var s1 scoreboard.Scoreboard
		s2 := scoreboard.Scoreboard{"ola": &scoreboard.Result{Name: "ola", Score: 5}}
		pOlof := player.Player{Name: "olof", Guess: &[2]int{0, 5}}
		err := s2.AddResult(2, pOlof)
		require.Nil(t, err, "error is not nil")
		s1 = scoreboard.Join(nil, s2)
		require.Equal(t, 0, s1[pOlof.Name].Score)
	})
}
