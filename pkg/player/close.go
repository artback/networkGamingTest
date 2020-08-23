package player

func (p *Player) Close() error{
	if p.ws != nil {
		return p.ws.Close()
	}else {
		return WebSocketNilError
	}
}
