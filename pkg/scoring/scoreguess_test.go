package scoring_test

import (
	"github.com/artback/networkGamingTest/pkg/scoring"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_scoreGuess(t *testing.T) {
	t.Run("Scoreguess with exact match, should give 5 points", func(t *testing.T) {
		score := scoring.ScoreGuess(4, [2]int{4, 6})
		require.Equal(t, 5, score)
	})
	t.Run("Scoreguess with two exact guess, should give 10 points", func(t *testing.T) {
		score := scoring.ScoreGuess(4, [2]int{4, 4})
		require.Equal(t, 10, score)
	})
	t.Run("Scoreguess with in of bounds, should give 4 points", func(t *testing.T) {
		score := scoring.ScoreGuess(3, [2]int{2, 6})
		require.Equal(t, 1, score)
	})
	t.Run("Scoreguess with out of bounds, should give -1 point", func(t *testing.T) {
		score := scoring.ScoreGuess(1, [2]int{2, 6})
		require.Equal(t, -1, score)
	})
	t.Run("Scoreguess with unordered guess, should give 5 points", func(t *testing.T) {
		score := scoring.ScoreGuess(4, [2]int{6, 4})
		require.Equal(t, 5, score)
	})
}
