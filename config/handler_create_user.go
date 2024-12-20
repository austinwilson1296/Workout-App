package config

import (
    "net/http"
    "encoding/json"
	"database/sql"
    "github.com/austinwilson1296/fitted/internal/database"
	"github.com/austinwilson1296/fitted/internal/auth"
	
)


func (cfg *ApiCfg) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
    type parameters struct {
        Email string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
		ExperienceLevelID int32 `json:"experience_level_id"`
    }
	type response struct{
		User
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Password == ""{
		respondWithError(w, http.StatusBadRequest, "password required", nil)
		return
	}

	if params.ExperienceLevelID == 0{
		params.ExperienceLevelID = 1
	}
	var experienceLevelSQL sql.NullInt32
	experienceLevelSQL.Int32 = params.ExperienceLevelID
	experienceLevelSQL.Valid = true

	hashedPassword, err:= auth.HashPassword(params.Password)
	if err != nil{
		respondWithError(w, http.StatusBadRequest, "could not set password",err)
		return
	}

	user,err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		PasswordHash: hashedPassword,
		Username: params.Username,
		Email: params.Email,
		ExperienceLevelID: experienceLevelSQL,
	})
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "could not create user", err)
		return
	}
	respondWithJSON(w,http.StatusCreated,response{
		User:User{
			ID: user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email: user.Email,
			Username: user.Username,
			ExperienceLevelID: user.ExperienceLevelID,
		},
	})
}
