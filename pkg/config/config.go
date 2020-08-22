package config

import (
	"github.com/artback/networkGamingTest/pkg/util/webservice"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
)

type Configuration struct {
	Rounds        int    ` env:"GAME_ROUNDS" env-default:"10" json:"rounds"`
	Sleep         string `env:"GAME_SLEEP_TIME" env-default:"1s" json:"sleep"`
	SleepBetween  string `env:"GAME_BETWEEN_SLEEP_TIME" env-default:"10s" json:"sleep_between"`
	MinimumPlayer int    `env:"GAME_MINIMUM_PLAYERS" env-default:"2" json:"minimum_player"`
	BeginInterval int    `env:"GAME_BEGIN_INTERVAL" env-default:"0" json:"begin_interval"`
	EndInterval   int    `env:"GAME_END_INTERVAL" env-default:"10" json:"end_interval"`
}

func (co *Configuration) Init() error {
	return cleanenv.ReadEnv(co)
}

func NewConfig() *Configuration {
	conf := &Configuration{}
	conf.Init()
	return conf
}

func (co *Configuration) ApiConfig(w http.ResponseWriter, _ *http.Request) {
	_, err := webservice.RespondWithJSON(w, http.StatusOK, co)
	if err != nil {
		log.Print(err)
	}
}
