package player

import (
	"github.com/gorilla/websocket"
	"time"
)

func (p *Player) Ping() error {
	if p.ws != nil {
		if err := p.ws.WriteControl(
			websocket.PingMessage,
			[]byte{},
			time.Now().Add(10*time.Millisecond),
		); err != nil {
			return err
		}
	} else {
		return WebSocketNilError
	}
	return nil
}
