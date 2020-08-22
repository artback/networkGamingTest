package scoreboard

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScoreboard_Winner(t *testing.T) {
	t.Run("get winner by point", func(t *testing.T) {
		s := Scoreboard{
			"ola":  &Result{Name: "ola", Score: 7, guess: &[2]int{1, 1}},
			"per":  &Result{Name: "per", Score: 1},
			"sven": &Result{Name: "sven", Score: 5, guess: &[2]int{5, 9}}}
		winner := s.Winner()
		require.Equal(t, 7, winner.Score)
		require.Equal(t, "ola", winner.Name)
	})
	t.Run("get winner by upperbound", func(t *testing.T) {
		s := Scoreboard{
			"ola":  &Result{Name: "ola", Score: 5, guess: &[2]int{1, 10}},
			"per":  &Result{Name: "per", Score: 1},
			"sven": &Result{Name: "sven", Score: 5, guess: &[2]int{5, 9}}}
		winner := s.Winner()
		require.Equal(t, 5, winner.Score)
		require.Equal(t, "ola", winner.Name)
	})
	t.Run("get winner by lower bound", func(t *testing.T) {
		s := Scoreboard{
			"ola":  &Result{Name: "ola", Score: 5, guess: &[2]int{1, 10}},
			"per":  &Result{Name: "per", Score: 1},
			"sven": &Result{Name: "sven", Score: 5, guess: &[2]int{5, 10}}}
		winner := s.Winner()
		require.Equal(t, 5, winner.Score)
		require.Equal(t, "sven", winner.Name)
	})
	t.Run("get winner by name ", func(t *testing.T) {
		s := Scoreboard{
			"ola":  &Result{Name: "ola", Score: 5},
			"per":  &Result{Name: "per", Score: 1},
			"sven": &Result{Name: "sven", Score: 5}}
		winner := s.Winner()
		require.Equal(t, &Result{Name: "ola", Score: 5}, winner)
	})
}
