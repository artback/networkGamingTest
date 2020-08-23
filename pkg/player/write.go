package player

import (
	"errors"
	"github.com/artback/networkGamingTest/internal/socketMessage"
)

func (p *Player) WriteMessage(Type string, payload interface{}) error {
	if p.ws != nil {
		return p.ws.WriteJSON(socketMessage.SocketMessage{Type: Type, Payload: payload})
	} else {
		return errors.New("socket is nil")
	}
}
