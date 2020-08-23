package gameService

import "log"

func (gs *GameService) PingPlayers() {
	for name, p := range gs.players {
		if p.Ping() != nil {
			log.Print("User ", name, " not answering ping")
			delete(gs.players, name)
			p.Close()
		}
	}
	for name, o := range gs.observers {
		if o.Ping() != nil {
			log.Print("User ", name, " not answering ping")
			delete(gs.observers, name)
			o.Close()
		}
	}
}
