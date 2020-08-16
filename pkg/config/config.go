package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	Game GameConfiguration
}
type GameConfiguration struct {
	Rounds   int    ` env:"GAME_ROUNDS"`
	Sleep    string `env:"GAME_SLEEP_TIME"`
	Interval [2]int `env:"GAME_INTERVAL"`
}

func Init(co *Configuration) error {
	return cleanenv.ReadEnv(co)
}
