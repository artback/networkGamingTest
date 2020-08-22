package player

import (
	"errors"
	"github.com/artback/networkGamingTest/internal/jsonwriter"
)

var WebSocketNilError = errors.New("websocket is nil")

type Player struct {
	Name  string
	ws    jsonwriter.JsonWriter
	Guess *[2]int
}

func NewPlayer(name string, ws jsonwriter.JsonWriter, guess *[2]int) Player {
	return Player{Name: name, ws: ws, Guess: guess}
}
