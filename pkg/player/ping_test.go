package player

import (
	"errors"
	"github.com/artback/networkGamingTest/internal/jsonwriter"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_Ping(t *testing.T) {
	t.Run("Ping socket timeout", func(t *testing.T) {
		mock := jsonwriter.Mock{Err: errors.New("timeout error")}
		p := NewPlayer("ola", mock, &[2]int{1, 10})
		err := p.Ping()
		require.Error(t, err, "timeout error")
	})
	t.Run("Ping socket ,return nil", func(t *testing.T) {
		mock := jsonwriter.Mock{}
		p := NewPlayer("ola", mock, &[2]int{1, 10})
		err := p.Ping()
		require.Nil(t, err)
	})
	t.Run("Ping socket is nil", func(t *testing.T) {
		p := NewPlayer("ola", nil, &[2]int{1, 10})
		err := p.Ping()
		require.Error(t, err, WebSocketNilError)
	})
}
