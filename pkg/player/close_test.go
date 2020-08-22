package player

import (
	"errors"
	"github.com/artback/networkGamingTest/internal/jsonwriter"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_Close(t *testing.T) {
	t.Run("Close, socket is nil", func(t *testing.T) {
		p := NewPlayer("ola", nil, &[2]int{1, 10})
		err := p.Close()
		require.Error(t, err, WebSocketNilError)
	})
	t.Run("Close, socket timeout", func(t *testing.T) {
		mock := jsonwriter.Mock{Err: errors.New("timeout error")}
		p := NewPlayer("ola", mock, &[2]int{1, 10})
		err := p.Close()
		require.Error(t, err, "timeout error")
	})
	t.Run("Close, return nil", func(t *testing.T) {
		mock := jsonwriter.Mock{}
		p := NewPlayer("ola", mock, &[2]int{1, 10})
		err := p.Close()
		require.Nil(t, err)
	})
}
