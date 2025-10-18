package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Key string
}

func NewConfig(name string) *Config {
	_ = godotenv.Load(name)
	return &Config{
		Key: os.Getenv("KEY"),
	}
}
