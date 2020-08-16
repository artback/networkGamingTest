package gamehandler

import (
	"fmt"
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/artback/networkGamingTest/pkg/game"
	"time"
)

func StartGame(c config.GameConfiguration, seed int64) error {
	duration, err := time.ParseDuration(c.Sleep)
	if err == nil {
		return err
	}
	game := game.NewGame(c.Rounds, seed, duration, c.Interval)
	resChan := make(chan int)
	go game.Run(resChan)

	// Do something with result
	for result := range resChan {
		fmt.Print(result)
	}
	return nil
}
