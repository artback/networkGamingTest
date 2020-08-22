package player

import (
	"errors"
)

func (p *Player) Ping() error {
	if p.ws != nil {
		if _, _, err := p.ws.NextReader(); err != nil {
			return err
		}
	} else {
		return errors.New("websocket is nil")
	}
	return nil
}
