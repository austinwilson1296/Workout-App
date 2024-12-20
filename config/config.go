package config

import (
    "github.com/austinwilson1296/fitted/internal/database"
)

type ApiCfg struct {
    Platform  string
    JwtSecret string
    DB        *database.Queries
}