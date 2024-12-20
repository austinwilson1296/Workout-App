package config

import (
    "net/http"
    "encoding/json"
	"github.com/austinwilson1296/fitted/internal/auth"
)

func (cfg *ApiCfg) HandlerLogin (w http.ResponseWriter, r *http.Request){
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
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
	user,err := cfg.DB.GetUserByUsername(r.Context(),params.Username)
	if err != nil{
		respondWithError(w,http.StatusBadRequest,"unable to locate user",err)
		return
	}
	err = auth.CheckPasswordHash(params.Password,user.PasswordHash)
	if err != nil{
		respondWithError(w, http.StatusUnauthorized,"invalid password",err)
		return
	}
	respondWithJSON(w, http.StatusOK,response{
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