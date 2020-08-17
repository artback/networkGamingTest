package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	Game GameConfiguration
}
type GameConfiguration struct {
	Rounds        int    ` env:"GAME_ROUNDS"`
	Sleep         string `env:"GAME_SLEEP_TIME" env-default:"1s"`
	SleepBetween  string `env:"GAME_BETWEEN_SLEEP_TIME" env-default:"10s"`
	Interval      [2]int `env:"GAME_INTERVAL"`
	MinimumPlayer int    `env:"GAME_MINIMUM_PLAYERS env-default:"2"`
}

func (co *Configuration) Init() error {
	return cleanenv.ReadEnv(co)
}

func NewConfig() *Configuration {
	conf := &Configuration{}
	conf.Init()
	return conf
}
