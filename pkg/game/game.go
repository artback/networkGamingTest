package game

import (
	"errors"
	"math/rand"
	"time"
)

type Settings struct {
	Rounds   int
	Seed     int64
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
	r := rand.New(rand.NewSource(c.Seed))
	for i := 0; i < c.Rounds; i++ {
		value := r.Intn(c.Interval[1]) + c.Interval[0]
		resChan <- value
		time.Sleep(c.Sleep)
	}
}

func NewGame(rounds int, seed int64, sleep time.Duration, interval [2]int) Settings {
	return Settings{Rounds: rounds, Seed: seed, Sleep: sleep, Interval: interval}
}
