package config

import (
    "net/http"
    "github.com/austinwilson1296/fitted/internal/database"
    "strconv"
    "sort"
)

func (cfg *ApiCfg) HandlerGenerateCoolDown(w http.ResponseWriter, r *http.Request) {
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
    coreHipsLegs, err := cfg.DB.GetCoreHipsLegsExercises(r.Context(), database.GetCoreHipsLegsExercisesParams{
        LevelID: levelInt,
        Limit:   4,
        Column3: true,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching core hips legs Exercises", err)
        return
    }

    coreSpinal, err := cfg.DB.GetCoreSpinalExercises(r.Context(), database.GetCoreSpinalExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching core spinal Exercises", err)
        return
    }

    thoracicSpine, err := cfg.DB.GetThoracicSpineMobilityExercises(r.Context(), database.GetThoracicSpineMobilityExercisesParams{
        LevelID: levelInt,
        Limit:   3,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching thoracic spine Exercises", err)
        return
    }

    scapuloThoracic, err := cfg.DB.GetScapuloThoracicExercises(r.Context(), database.GetScapuloThoracicExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching scapulo thoracic Exercises", err)
        return
    }

    shouldersScapula, err := cfg.DB.GetShouldersScapulaExercises(r.Context(), database.GetShouldersScapulaExercisesParams{
        LevelID: levelInt,
        Limit:   limit,
    })
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error fetching shoulders scapula Exercises", err)
        return
    }

    // Combine all Exercises
    returnWarmUp := []Exercise{}
    
    coreHipLegPositions := []int{0, 3, 7, 10}
    for i, place := range coreHipLegPositions {
        if i < len(coreHipsLegs) { 
            returnWarmUp = append(returnWarmUp, Exercise{
                ExerciseID:   coreHipsLegs[i].ExerciseID,
                ExerciseName: coreHipsLegs[i].ExerciseName,
                CategoryName: coreHipsLegs[i].CategoryName,
                ListPosition: place,
            })
        }
    }

    
    coreSpinalPositions := []int{1, 4}
    for i, place := range coreSpinalPositions {
        if i < len(coreSpinal) {
            returnWarmUp = append(returnWarmUp, Exercise{
                ExerciseID:   coreSpinal[i].ExerciseID,
                ExerciseName: coreSpinal[i].ExerciseName,
                CategoryName: coreSpinal[i].CategoryName,
                ListPosition: place,
            })
        }
    }
    thoracicSpinePositions := []int{2,5,9}
    for i,place := range thoracicSpinePositions{
        if i < len(thoracicSpine){
        returnWarmUp = append(returnWarmUp, Exercise{
            ExerciseID:   thoracicSpine[i].ExerciseID,
            ExerciseName: thoracicSpine[i].ExerciseName,
            CategoryName: thoracicSpine[i].CategoryName,
            ListPosition: place,
        })
    }
    }
    
    scapuloThoracicPositions := []int{6}
    for i,place := range scapuloThoracicPositions{
        if i < len(scapuloThoracic){
        returnWarmUp = append(returnWarmUp, Exercise{
            ExerciseID:   scapuloThoracic[i].ExerciseID,
            ExerciseName: scapuloThoracic[i].ExerciseName,
            CategoryName: scapuloThoracic[i].CategoryName,
            ListPosition: place,
        })
    }
    }
    
    shouldersScapulaPositions := []int{8,11}
    for i,place := range shouldersScapulaPositions{
        if i < len(shouldersScapula){
        returnWarmUp = append(returnWarmUp, Exercise{
            ExerciseID:   shouldersScapula[i].ExerciseID,
            ExerciseName: shouldersScapula[i].ExerciseName,
            CategoryName: shouldersScapula[i].CategoryName,
            ListPosition: place,
        })
    }
    }
    sort.Slice(returnWarmUp, func(i, j int) bool {
        return returnWarmUp[i].ListPosition < returnWarmUp[j].ListPosition
    })

    // Respond with JSON data
    respondWithJSON(w, http.StatusOK, returnWarmUp)
}
