package gameService

import (
	"github.com/artback/networkGamingTest/pkg/util/webservice"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

// This method contains everything necessary to join a game and setup a websocket connection
func (gs *gameService) ApiGameWs(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// Something along these lines to upgrade to a Websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		webservice.RespondWithError(w, r, http.StatusBadRequest, "'upgrade' token not found in 'Connection' header")
		return
	}

	qGuess := r.URL.Query()["guess"]
	var guess [2]int
	guess[0], _ = strconv.Atoi(qGuess[0])
	guess[1], _ = strconv.Atoi(qGuess[1])

	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		webservice.RespondWithError(w, r, http.StatusBadRequest, "name required")
		return
	}

	gs.Join(name, ws, guess)
}
