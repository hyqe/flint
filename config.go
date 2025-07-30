package main

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type config struct {
	Server serverConfig `envPrefix:"SERVER_"`
}

func newConfig() config {
	return config{}
}

func parseConfig(c *config) error {
	_ = godotenv.Load()
	if err := env.Parse(c); err != nil {
		return fmt.Errorf("failed to parse env: %w", err)
	}
	return nil
}
