package config

import (
    "github.com/austinwilson1296/fitted/internal/database"
    "github.com/google/uuid"
    "time"
    "database/sql"
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

type User struct{
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Username string `json:"username"`
	ExperienceLevelID sql.NullInt32 `json:"experience_level_id"`
}