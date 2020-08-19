package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	Game GameConfiguration
}
type GameConfiguration struct {
	Rounds        int    ` env:"GAME_ROUNDS" env-default:"1"`
	Sleep         string `env:"GAME_SLEEP_TIME" env-default:"1s"`
	SleepBetween  string `env:"GAME_BETWEEN_SLEEP_TIME" env-default:"10s"`
	MinimumPlayer int    `env:"GAME_MINIMUM_PLAYERS" env-default:"2"`
	BeginInterval int    `env:"GAME_BEGIN_INTERVAL" env-default:"0"`
	EndInterval   int    `env:"GAME_END_INTERVAL" env-default:"10"`
}

func (co *Configuration) Init() error {
	return cleanenv.ReadEnv(co)
}

func NewConfig() *Configuration {
	conf := &Configuration{}
	conf.Init()
	return conf
}
