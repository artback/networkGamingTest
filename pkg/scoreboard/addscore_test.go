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
		err := s.AddResult(5, p)
		guess := [2]int{0, 1}
		name := "ola"
		require.Equal(t, guess, *s[name].guess)
		require.Equal(t, -1, s[name].Score)
		require.Equal(t, name, s[name].Name)
		require.Nil(t, err)
	})
	t.Run("guess is nil", func(t *testing.T) {
		s := Scoreboard{}
		p := player.NewPlayer("ola", nil, nil)
		err := s.AddResult(5, p)
		require.EqualError(t, err, "guess is nil")
	})
}

func TestScoreboard_AddScore(t *testing.T) {
	t.Run("Join scores on user from two tables", func(t *testing.T) {
		s1 := Scoreboard{}
		s2 := Scoreboard{}
		pOlof := player.Player{Name: "olof", Guess: &[2]int{0, 5}}
		pBjorn := player.Player{Name: "bjorn", Guess: &[2]int{0, 2}}
		err := s1.AddResult(2, pOlof)
		require.Nil(t, err, "error is not nil")
		err = s1.AddResult(10, pBjorn)
		require.Nil(t, err, "error is not nil")
		err = s2.AddResult(2, pOlof)
		require.Nil(t, err, "error is not nil")
		s1 = Join(s1, s2)
		require.Equal(t, 0, s1[pOlof.Name].Score)
		require.Equal(t, "olof", s1[pOlof.Name].Name)
	})
}
