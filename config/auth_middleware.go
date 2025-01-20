package config

import (
	"net/http"
	"github.com/austinwilson1296/fitted/internal/auth"
)

func (cfg *ApiCfg) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt-token")
		if err != nil {
			// respondWithError(w, http.StatusUnauthorized, "No valid session found", nil)
			http.Redirect(w,r,"/login",http.StatusSeeOther)
			return
		}

		
		userID, err := auth.ValidateJWT(cookie.Value, cfg.JwtSecret)
		
		
		if err != nil {
			http.Redirect(w,r,"/login",http.StatusSeeOther)
		}

		ctx := auth.ContextWithUserID(r.Context(), userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}