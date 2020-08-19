package player

import (
	"errors"
	"github.com/artback/networkGamingTest/internal/jsonwriter"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_AddScore(t *testing.T) {
	t.Run("Addscore to player", func(t *testing.T) {
		s := scoreboard.NewScoreBoard(1)
		mock := jsonwriter.Mock{Err: nil}
		p := NewPlayer("ola", mock, [2]int{0, 1})
		_ = p.AddScore(5, s)
		require.Equal(t, map[string]int{"ola": -1}, s.Map())
	})
	t.Run("Addscore to player", func(t *testing.T) {
		wError := errors.New("write error")
		s := scoreboard.NewScoreBoard(1)
		mock := jsonwriter.Mock{Err: wError}
		p := NewPlayer("ola", mock, [2]int{0, 1})
		err := p.AddScore(5, s)
		require.EqualError(t, err, wError.Error())
	})
}
