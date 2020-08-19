package player

import (
	"github.com/artback/networkGamingTest/internal/jsonwriter"
)

type Player struct {
	Name  string
	Jw    jsonwriter.JsonWriter
	Guess [2]int
}

func NewPlayer(name string, jw jsonwriter.JsonWriter, guess [2]int) Player {
	return Player{Name: name, Jw: jw, Guess: guess}
}
