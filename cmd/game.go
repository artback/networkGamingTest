package main

import (
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/artback/networkGamingTest/pkg/gameService"
	"github.com/artback/networkGamingTest/pkg/scoreboard"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	c := config.NewConfig()
	total := scoreboard.NewScoreBoard(10)
	game := gameService.NewGameService()
	go game.GameLoop(c.Game, time.Now().UnixNano(), total)
	r := mux.NewRouter()
	r.HandleFunc("/api/game/ws", game.ApiGameWs)
	log.Print(http.ListenAndServe(":8080", r))
}
