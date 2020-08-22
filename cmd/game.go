package main

import (
	"github.com/artback/networkGamingTest/pkg/api"
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/artback/networkGamingTest/pkg/gameService"
	"time"
)

func main() {
	c := config.NewConfig()
	gs := gameService.NewGameService()
	go gameService.GameLoop(gs, *c, time.Now().UnixNano())
	h := api.NewHandler(c, gs)
	h.Api()
}
