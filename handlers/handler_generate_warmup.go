package handler

import (
	"net/http"
	"encoding/json"
	"github.com/austinwilson1296/fitted/internal/database"
)

type warmUp struct {
	ExerciseID   int32  `json:"id"`
	ExerciseName string `json:"exercise name"`
	CategoryName string `json:"category name"`
}

func HandlerGenerateWarmUp(cfg *ApiCfg, w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Level int32 `json:"level"`
	}

	var limit int32 = 2

	// Decode parameters from the request body
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	// No need to convert level since it's already an integer
	levelInt := params.Level

	// Fetch exercises from the database
	dbWarmUp, err := cfg.db.GetExercisesPerCategoryByExperience(r.Context(), database.GetExercisesPerCategoryByExperienceParams{
		LevelID: levelInt,
		Limit:   limit,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error fetching workout", err)
		return
	}

	// Prepare the return data
	returnWarmUp := []warmUp{}
	for _, item := range dbWarmUp {
		returnWarmUp = append(returnWarmUp, warmUp{
			ExerciseID:   item.ExerciseID,
			ExerciseName: item.ExerciseName,
			CategoryName: item.CategoryName,
		})
	}

	// Respond with JSON data
	respondWithJSON(w, http.StatusOK, returnWarmUp)
}