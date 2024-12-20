package config

import (
    "github.com/austinwilson1296/fitted/internal/database"
)

type ApiCfg struct {
    Platform  string
    JwtSecret string
    DB        *database.Queries
}

type Exercise struct {
    ExerciseID   int32  `json:"id"`
    ExerciseName string `json:"exercise_name"`
    CategoryName string `json:"category_name"`
}