package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Run("default configuration ", func(t *testing.T) {
		c := NewConfig()
		expected := Configuration{
			Rounds:        10,
			Sleep:         "1s",
			SleepBetween:  "10s",
			BeginInterval: 0,
			EndInterval:   10,
			MinimumPlayer: 2,
		}
		require.Equal(t, expected, *c)
	})
}
