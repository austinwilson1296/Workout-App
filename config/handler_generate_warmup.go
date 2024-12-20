package config

import (
    "net/http"
    "encoding/json"
    "github.com/austinwilson1296/fitted/internal/database"
)

type warmUp struct {
    ExerciseID   int32  `json:"id"`
    ExerciseName string `json:"exercise_name"`
    CategoryName string `json:"category_name"`
}

func (cfg *ApiCfg) HandlerGenerateWarmUp(w http.ResponseWriter, r *http.Request) {
    type parameters struct {
        Level int32 `json:"level"`
    }

    var limit int32 = 2

    decoder := json.NewDecoder(r.Body)
    params := parameters{}
    err := decoder.Decode(&params)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
        return
    }

    levelInt := params.Level
    ctx := r.Context()

    // Get exercises for each category
    coreHipsLegs, err := cfg.DB.GetCoreHipsLegsExercises(ctx, database.GetCoreHipsLegsExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching core hips legs exercises", err)
        return
    }

    coreSpinal, err := cfg.DB.GetCoreSpinalExercises(ctx, database.GetCoreSpinalExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching core spinal exercises", err)
        return
    }

    thoracicSpine, err := cfg.DB.GetThoracicSpineMobilityExercises(ctx, database.GetThoracicSpineMobilityExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching thoracic spine exercises", err)
        return
    }

    scapuloThoracic, err := cfg.DB.GetScapuloThoracicExercises(ctx, database.GetScapuloThoracicExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching scapulo thoracic exercises", err)
        return
    }

    shouldersScapula, err := cfg.DB.GetShouldersScapulaExercises(ctx, database.GetShouldersScapulaExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching shoulders scapula exercises", err)
        return
    }

    // Combine all exercises
    returnWarmUp := []warmUp{}
    
    // Add exercises from each category
    for _, item := range coreHipsLegs {
        returnWarmUp = append(returnWarmUp, warmUp{
            ExerciseID:   item.ExerciseID,
            ExerciseName: item.ExerciseName,
            CategoryName: item.CategoryName,
        })
    }
    
    for _, item := range coreSpinal {
        returnWarmUp = append(returnWarmUp, warmUp{
            ExerciseID:   item.ExerciseID,
            ExerciseName: item.ExerciseName,
            CategoryName: item.CategoryName,
        })
    }
    
    for _, item := range thoracicSpine {
        returnWarmUp = append(returnWarmUp, warmUp{
            ExerciseID:   item.ExerciseID,
            ExerciseName: item.ExerciseName,
            CategoryName: item.CategoryName,
        })
    }
    
    for _, item := range scapuloThoracic {
        returnWarmUp = append(returnWarmUp, warmUp{
            ExerciseID:   item.ExerciseID,
            ExerciseName: item.ExerciseName,
            CategoryName: item.CategoryName,
        })
    }
    
    for _, item := range shouldersScapula {
        returnWarmUp = append(returnWarmUp, warmUp{
            ExerciseID:   item.ExerciseID,
            ExerciseName: item.ExerciseName,
            CategoryName: item.CategoryName,
        })
    }

    // Respond with JSON data
    respondWithJSON(w, http.StatusOK, returnWarmUp)
}