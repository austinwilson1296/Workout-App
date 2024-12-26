package config

import (
	"fmt"
	"net/http"
	"github.com/austinwilson1296/fitted/internal/auth"
)

func (cfg *ApiCfg) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt-token")
		if err != nil {
			fmt.Printf("Cookie error: %v\n", err)
			respondWithError(w, http.StatusUnauthorized, "No valid session found", nil)
			return
		}

		fmt.Printf("Cookie present: %v\n", cookie != nil)
		fmt.Printf("Cookie value: %v\n", cookie.Value)

		userID, err := auth.ValidateJWT(cookie.Value, cfg.JwtSecret)
		fmt.Printf("JWT validation error: %v\n", err)
		
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Invalid session", nil)
			return
		}

		ctx := auth.ContextWithUserID(r.Context(), userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}