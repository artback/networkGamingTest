package gameService

import (
	"github.com/artback/networkGamingTest/internal/socketMessage"
	"github.com/artback/networkGamingTest/pkg/util/webservice"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sort"
	"strconv"
)

// This method contains everything necessary to join a game and setup a websocket connection
func (gs *GameService) ApiGameWs(w http.ResponseWriter, r *http.Request) {
	qGuess := r.URL.Query()["guess"]
	var guess *[2]int
	if len(qGuess) > 0 {
		g1, err := strconv.Atoi(qGuess[0])
		g2, err := strconv.Atoi(qGuess[1])
		if err == nil {
			guess = &[2]int{g1, g2}
		}
	}
	sort.Slice(guess, func(i, j int) bool {
		return guess[i] > guess[j]
	})

	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		_, err := webservice.RespondWithError(w, r, http.StatusBadRequest, "name required")
		if err != nil {
			log.Print("ApiGameWS ", err)
		}
		return
	}
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// Something along these lines to upgrade to a Websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if ws != nil {
		if err != nil {
			err := ws.WriteJSON(socketMessage.SocketMessage{Type: "error", Payload: "'upgrade' token not found in 'Connection' header"})
			if err != nil {
				log.Print("ApiGameWS websocket", err)
			}
			err = ws.Close()
			if err != nil {
				log.Print("ApiGameWS websocket ", err)
			}
			return
		}
		err = ws.WriteJSON(socketMessage.SocketMessage{Type: "total", Payload: gs.total})
		if err != nil {
			log.Print("ApiGameWS websocket ", err)
		}
	}
	gs.Join(name, ws, guess)

}
