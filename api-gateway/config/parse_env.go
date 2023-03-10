package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

func (c *Config) ParseEnv() error {
	envHTTPServHost := os.Getenv(envNameHTTPServHost)
	if envHTTPServHost != "" {
		c.HTTPServHost = envHTTPServHost
	}

	envHTTPServPort := os.Getenv(envNameHTTPServPort)
	if envHTTPServPort != "" {
		c.HTTPServPort = envHTTPServPort
	}

	envLogLvl := os.Getenv(envNameLogLvl)
	if envLogLvl != "" {
		logLevel, err := zerolog.ParseLevel(envLogLvl)
		if err != nil {
			return fmt.Errorf("parse log lvl: %s", envLogLvl)
		}
		c.LogLvl = logLevel
	}

	return nil
}
