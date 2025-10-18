package api

import (
	"app-3/config"
)

func GetKey(cfg *config.Config) string {
	return cfg.Key
}
