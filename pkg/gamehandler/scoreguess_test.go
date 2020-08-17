package gamehandler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_scoreGuess(t *testing.T) {
	t.Run("Scoreguess with exact match, should give 5 points", func(t *testing.T) {
		score := scoreGuess(4, [2]int{4, 6})
		require.Equal(t, 5, score)
	})
	t.Run("Scoreguess with in of bounds, should give 4 points", func(t *testing.T) {
		score := scoreGuess(3, [2]int{2, 6})
		require.Equal(t, 1, score)
	})
	t.Run("Scoreguess with out of bounds, should give -1 point", func(t *testing.T) {
		score := scoreGuess(1, [2]int{2, 6})
		require.Equal(t, -1, score)
	})
}
