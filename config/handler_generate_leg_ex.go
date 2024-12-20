package config 

import (
    "net/http"
	"strconv"
)

func (cfg *ApiCfg) HandlerGenerateLegExercise(w http.ResponseWriter, r *http.Request) {
    levelStr := r.URL.Query().Get("level")
    level, err := strconv.ParseInt(levelStr, 10, 32)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid level parameter", err)
        return
    }
	type response struct{
		Exercise
	}
    
    legExercise, err := cfg.DB.GetLegExercises(r.Context(), int32(level))
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "could not locate exercise", err)
        return
    }
    
    respondWithJSON(w, http.StatusOK, response{
        Exercise: Exercise{
            ExerciseID:   legExercise.ExerciseID,
            ExerciseName: legExercise.ExerciseName,
            CategoryName: legExercise.CategoryName,
        },
    })
}