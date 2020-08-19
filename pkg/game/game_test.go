package game

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func Test_Run(t *testing.T) {
	t.Run("seed 5,interval 0,10", func(t *testing.T) {
		g := Settings{
			Seed:     rand.NewSource(5),
			Sleep:    0,
			Rounds:   2,
			Interval: [2]int{0, 10},
		}
		expected := []int{6, 6}
		resChan := make(chan int)
		errorChan := make(chan error)
		go g.Run(resChan, errorChan)
		var actual []int
		for result := range resChan {
			actual = append(actual, result)
		}
		require.Equal(t, expected, actual)
	})
	t.Run("interval error", func(t *testing.T) {
		g := Settings{
			Seed:     rand.NewSource(5),
			Sleep:    0,
			Rounds:   2,
			Interval: [2]int{0, 0},
		}
		resChan := make(chan int)
		errorChan := make(chan error)
		go g.Run(resChan, errorChan)
		var actual []int
		err := <-errorChan
		for result := range resChan {
			actual = append(actual, result)
		}
		var expected []int
		require.EqualError(t, err, "interval is undefined")
		require.Equal(t, expected, actual)
	})
}
