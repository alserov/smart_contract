package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	ContractAddr string
}

func MustLoad() *Config {
	var cfg Config

	cfg.ContractAddr = os.Getenv("CONTRACT_ADDR")

	return &cfg
}
