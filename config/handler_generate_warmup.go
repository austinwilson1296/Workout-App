package config

import (
    "net/http"
    "github.com/austinwilson1296/fitted/internal/database"
    "strconv"
    "github.com/austinwilson1296/fitted/internal/auth"
    "context"
)

func (cfg *ApiCfg) HandlerGenerateWarmUp(w http.ResponseWriter, r *http.Request) {
    claims, err := ParseAndValidateToken(w, r)
    if err != nil {
        return 
    }

    // Get level from query parameter
    levelStr := r.URL.Query().Get("level")
    level, err := strconv.ParseInt(levelStr, 10, 32)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid level parameter", err)
        return
    }
    var limit int32 = 2
    levelInt := int32(level)

    // Get Exercises for each category
    coreHipsLegs, err := cfg.DB.GetCoreHipsLegsExercises(ctx, database.GetCoreHipsLegsExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching core hips legs Exercises", err)
        return
    }

    coreSpinal, err := cfg.DB.GetCoreSpinalExercises(ctx, database.GetCoreSpinalExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching core spinal Exercises", err)
        return
    }

    thoracicSpine, err := cfg.DB.GetThoracicSpineMobilityExercises(ctx, database.GetThoracicSpineMobilityExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching thoracic spine Exercises", err)
        return
    }

    scapuloThoracic, err := cfg.DB.GetScapuloThoracicExercises(ctx, database.GetScapuloThoracicExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching scapulo thoracic Exercises", err)
        return
    }

    shouldersScapula, err := cfg.DB.GetShouldersScapulaExercises(ctx, database.GetShouldersScapulaExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching shoulders scapula Exercises", err)
        return
    }

    // Combine all Exercises
    returnWarmUp := []Exercise{}
    
    // Add Exercises from each category
    for _, item := range coreHipsLegs {
        returnWarmUp = append(returnWarmUp, Exercise{
            ExerciseID:   item.ExerciseID,
            ExerciseName: item.ExerciseName,
            CategoryName: item.CategoryName,
        })
    }
    
    for _, item := range coreSpinal {
        returnWarmUp = append(returnWarmUp, Exercise{
            ExerciseID:   item.ExerciseID,
            ExerciseName: item.ExerciseName,
            CategoryName: item.CategoryName,
        })
    }
    
    for _, item := range thoracicSpine {
        returnWarmUp = append(returnWarmUp, Exercise{
            ExerciseID:   item.ExerciseID,
            ExerciseName: item.ExerciseName,
            CategoryName: item.CategoryName,
        })
    }
    
    for _, item := range scapuloThoracic {
        returnWarmUp = append(returnWarmUp, Exercise{
            ExerciseID:   item.ExerciseID,
            ExerciseName: item.ExerciseName,
            CategoryName: item.CategoryName,
        })
    }
    
    for _, item := range shouldersScapula {
        returnWarmUp = append(returnWarmUp, Exercise{
            ExerciseID:   item.ExerciseID,
            ExerciseName: item.ExerciseName,
            CategoryName: item.CategoryName,
        })
    }

    // Respond with JSON data
    respondWithJSON(w, http.StatusOK, returnWarmUp)
}
