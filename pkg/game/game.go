package game

import (
	"errors"
	"math/rand"
	"time"
)

type Settings struct {
	Rounds   int
	Seed     rand.Source
	Sleep    time.Duration
	Interval [2]int
}

//Runs for Rounds and returns a randomized value in Interval.
//Sleeps for Sleep duration between rounds.
//Returns the result through resChan and errors through errorChan
func (c Settings) Run(resChan chan int, errorChan chan error) {
	defer close(resChan)
	defer close(errorChan)
	if c.Interval[1] <= 0 {
		errorChan <- errors.New("interval is undefined")
		return
	}
	r := rand.New(c.Seed)
	for i := 0; i < c.Rounds; i++ {
		value := r.Intn(c.Interval[1]) + c.Interval[0]
		resChan <- value
		time.Sleep(c.Sleep)
	}
}

func NewGame(rounds int, seed int64, sleep time.Duration, interval [2]int) Settings {
	source := rand.NewSource(seed)
	return Settings{Rounds: rounds, Seed: source, Sleep: sleep, Interval: interval}
}
