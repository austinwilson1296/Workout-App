package config

import (
	"net/http"
	"time"
	"github.com/austinwilson1296/fitted/internal/auth"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Every(time.Second), 3) // 3 requests per second

func (cfg *ApiCfg) HandlerLogin(w http.ResponseWriter, r *http.Request) {
    if !limiter.Allow() {
        respondWithError(w, http.StatusTooManyRequests, "Rate limit exceeded", nil)
        return
    }

    if err := r.ParseForm(); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid form data", err)
        return
    }

    username := r.Form.Get("username")
    password := r.Form.Get("password")

    if username == "" || password == "" {
        respondWithError(w, http.StatusBadRequest, "Username and password are required", nil)
        return
    }

    user, err := cfg.DB.GetUserByUsername(r.Context(), username)
    if err != nil {
        // Use same error message to prevent username enumeration
        respondWithError(w, http.StatusUnauthorized, "Invalid credentials", nil)
        return
    }

    if err := auth.CheckPasswordHash(password, user.PasswordHash); err != nil {
        respondWithError(w, http.StatusUnauthorized, "Invalid credentials", nil)
        return
    }

    accessToken, err := auth.MakeJWT(user.ID, cfg.JwtSecret, time.Hour)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Error creating session", nil)
        return
    }

    cookie := http.Cookie{
        Name:     "jwt-token",
        Value:    accessToken,
        Path:     "/",
        Secure:   true,
        HttpOnly: true,
        SameSite: http.SameSiteStrictMode,
        MaxAge:   2000, 
    }

    http.SetCookie(w, &cookie)
    w.Header().Set("HX-Redirect", "/workout")
    w.WriteHeader(http.StatusOK)
}

    
