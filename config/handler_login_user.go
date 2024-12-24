package config

import (
	"net/http"
	"fmt"
	// "encoding/json"
	"github.com/austinwilson1296/fitted/internal/auth"
)

func (cfg *ApiCfg) HandlerLogin(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Login attempt received") // Add logging

    if err := r.ParseForm(); err != nil {
        fmt.Println("Form parse error:", err) // Add logging
        respondWithError(w, http.StatusBadRequest, "Invalid form data", nil)
        return
    }

    username := r.Form.Get("username")
    password := r.Form.Get("password")
    fmt.Printf("Login attempt for username: %s\n", username) // Add logging

    if username == "" || password == "" {
        fmt.Println("Empty username or password") // Add logging
        respondWithError(w, http.StatusBadRequest, "Username and password are required", nil)
        return
    }

    user, err := cfg.DB.GetUserByUsername(r.Context(), username)
    if err != nil {
        fmt.Printf("Database error: %v\n", err) // Add logging
        respondWithError(w, http.StatusUnauthorized, "Invalid credentials", nil)
        return
    }

    if err := auth.CheckPasswordHash(password, user.PasswordHash); err != nil {
        fmt.Printf("Password check failed: %v\n", err) // Add logging
        respondWithError(w, http.StatusUnauthorized, "Invalid credentials", nil)
        return
    }

    fmt.Println("Login successful, redirecting") // Add logging
    w.Header().Set("HX-Redirect", "/workout")
    w.WriteHeader(http.StatusOK) // Respond with success
    
}