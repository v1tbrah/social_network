package config

import (
	"github.com/rs/zerolog"
)

type Config struct {
	HTTPServHost string
	HTTPServPort string
	LogLvl       zerolog.Level
}

func NewDefaultConfig() Config {
	return Config{
		HTTPServHost: defaultHTTPServHost,
		HTTPServPort: defaultHTTPServPort,
		LogLvl:       defaultLogLvl,
	}
}
