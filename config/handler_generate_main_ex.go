package config 

import (
    "net/http"
	"strconv"
    "github.com/austinwilson1296/fitted/internal/database"
)

func (cfg *ApiCfg) HandlerGenerateMainExercise(w http.ResponseWriter, r *http.Request) {
    // Get the query parameters
    levelStr := r.URL.Query().Get("level")
    level, err := strconv.ParseInt(levelStr, 10, 32)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid level parameter", err)
        return
    }
    name := r.URL.Query().Get("name")
    if name == "" {
        respondWithError(w, http.StatusBadRequest, "Missing required parameter: name", nil)
        return
    }

    // Prepare response structure
    type response struct {
		Exercise
	}
    
    // Get the exercise details from the database
    exercise, err := cfg.DB.GetMainExercise(r.Context(), database.GetMainExerciseParams{
        Name:    name,
        LevelID: int32(level),
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Could not locate exercise", err)
        return
    }
    
    // Return the exercise details in the response
    respondWithJSON(w, http.StatusOK, response{
        Exercise: Exercise{
            ExerciseID:   exercise.ExerciseID,
            ExerciseName: exercise.ExerciseName,
            CategoryName: exercise.CategoryName,
        },
    })
}
