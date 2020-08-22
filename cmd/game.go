package main

import (
	"github.com/artback/networkGamingTest/pkg/api"
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/artback/networkGamingTest/pkg/gameService"
	"log"
	"os"
	"time"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	gs := gameService.NewGameService()
	go gameService.GameLoop(gs, *c, time.Now().UnixNano())
	h := api.NewHandler(c, gs)
	if err := h.Api(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
