package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"social_network/api-gateway/config"
	api "social_network/api-gateway/internal"
)

func main() {
	newConfig := config.NewDefaultConfig()
	zerolog.SetGlobalLevel(newConfig.LogLvl)

	if err := newConfig.ParseEnv(); err != nil {
		log.Fatal().Err(err).Msg("config.ParseEnv")
	}
	zerolog.SetGlobalLevel(newConfig.LogLvl)

	newAPI := api.New(newConfig)

	newAPI.Run()
}
