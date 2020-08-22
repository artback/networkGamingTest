package api

import (
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/artback/networkGamingTest/pkg/gameService"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type Handler struct {
	config      *config.Configuration
	gameService *gameService.GameService
}

func (h *Handler) Api() {
	r := mux.NewRouter()
	r.HandleFunc("/api/game/ws", h.gameService.ApiGameWs)
	r.HandleFunc("/api/config", h.config.ApiConfig)
	handler := cors.Default().Handler(r)

	log.Print(http.ListenAndServe(":8080", handler))
}
func NewHandler(config *config.Configuration, gameService *gameService.GameService) *Handler {
	return &Handler{config: config, gameService: gameService}
}
