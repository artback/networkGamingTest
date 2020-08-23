package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Run("default configuration ", func(t *testing.T) {
		c, err := NewConfig()
		expected := Configuration{
			Rounds:        30,
			Sleep:         "1s",
			SleepBetween:  "10s",
			BeginInterval: 0,
			EndInterval:   10,
			MinimumPlayer: 2,
		}
		require.NotNil(t, c)
		if c != nil {
			require.Equal(t, expected, *c)
			require.Nil(t, err)
		}
	})
}
