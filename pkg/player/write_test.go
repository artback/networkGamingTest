package player

import (
	"errors"
	"github.com/artback/networkGamingTest/internal/jsonwriter"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlayer_WriteMessage(t *testing.T) {

	t.Run("WriteMessage, socket is nil", func(t *testing.T) {
		p := NewPlayer("ola", nil, &[2]int{1, 10})
		err := p.WriteMessage("type", []byte{})
		require.Error(t, err, WebSocketNilError)
	})
	t.Run("WriteMessage, socket timeout", func(t *testing.T) {
		mock := jsonwriter.Mock{Err: errors.New("timeout error")}
		p := NewPlayer("ola", mock, &[2]int{1, 10})
		err := p.WriteMessage("type", []byte{})
		require.Error(t, err, "timeout error")
	})
	t.Run("Close, return nil", func(t *testing.T) {
		mock := jsonwriter.Mock{}
		p := NewPlayer("ola", mock, &[2]int{1, 10})
		err := p.WriteMessage("type", []byte{})
		require.Nil(t, err)
	})
}
