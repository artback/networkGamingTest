package api

import (
	"github.com/artback/networkGamingTest/pkg/config"
	"github.com/artback/networkGamingTest/pkg/gameService"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	config      *config.Configuration
	gameService *gameService.GameService
}

func (h *Handler) Api() error {
	r := mux.NewRouter()
	r.HandleFunc("/ws/game", h.gameService.ApiGameWs)

	r.HandleFunc("/api/config", h.config.ApiConfig)

	return http.ListenAndServe(":8080", r)
}
func NewHandler(config *config.Configuration, gameService *gameService.GameService) *Handler {
	return &Handler{config: config, gameService: gameService}
}
