package scoreboard_test

import (
	"github.com/artback/networkGamingTest/pkg/player"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScoreboard_AddScore(t *testing.T) {
	t.Run("Join scores on user from two tables", func(t *testing.T) {
		s1 := scoreboard.Scoreboard{}
		s2 := scoreboard.Scoreboard{}
		pOlof := player.Player{Name: "bjorn", Guess: &[2]int{0, 5}}
		pBjorn := player.Player{Name: "bjorn", Guess: &[2]int{0, 2}}
		err := s1.AddResult(2, pOlof)
		require.Nil(t, err, "error is not nil")
		err = s1.AddResult(10, pBjorn)
		require.Nil(t, err, "error is not nil")
		err = s2.AddResult(2, pOlof)
		require.Nil(t, err, "error is not nil")
		s1 = scoreboard.Join(s1, s2)
		require.Equal(t, -1, s1[pOlof.Name])
		require.Equal(t, -1, s1[pBjorn.Name])
	})
}

func TestScoreboard_JoinScoreboards(t *testing.T) {
	t.Run("joining two scoreboards together", func(t *testing.T) {
		s1 := scoreboard.Scoreboard{}
		s2 := scoreboard.Scoreboard{}
		pOlof := player.Player{Name: "olof",Guess: &[2]int{0,1}}
		err := s1.AddResult(2, pOlof)
		require.Nil(t, err, "error is not nil")
		err = s2.AddResult(2, pOlof)
		require.Nil(t, err, "error is not nil")
		sJoin := scoreboard.Join(s1, s2)
		score := sJoin[pOlof.Name]
		require.Equal(t, -2, score)
	})
	t.Run("joining nil scoreboard with non nil", func(t *testing.T) {
		var s1 scoreboard.Scoreboard
		s2 := scoreboard.Scoreboard{"ola": 5}
		pOlof := player.Player{Name: "olof",Guess: &[2]int{0,1}}
		err :=  s2.AddResult(2, pOlof)
		require.Nil(t, err, "error is not nil")
		s1 = scoreboard.Join(nil, s2)
		score := s1[pOlof.Name]
		require.Equal(t, -1, score)
	})
	t.Run("joining two nil scoreboard ", func(t *testing.T) {
		var s1 scoreboard.Scoreboard
		s2 := scoreboard.Scoreboard{"ola": 5}
		pOlof := player.Player{Name: "olof",Guess: &[2]int{0,5}}
		err := s2.AddResult(2, pOlof)
		require.Nil(t, err, "error is not nil")
		s1 = scoreboard.Join(nil, s2)
		require.Equal(t, 0, s1[pOlof.Name])
	})
}
