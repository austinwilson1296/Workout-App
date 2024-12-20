package config

import (
    "errors"
    "net/http"
    "github.com/austinwilson1296/fitted/internal/auth"
   
)

// ParseAndValidateToken is a helper function to parse and validate the JWT token from the Authorization header.
// It returns the parsed claims and an error if any issues occur.
func ParseAndValidateToken(w http.ResponseWriter, r *http.Request) (*auth.Claims, error) {
    tokenString := r.Header.Get("Authorization")
    if tokenString == "" {
        respondWithError(w, http.StatusUnauthorized, "Missing token", nil)
        return nil, errors.New("missing token")
    }

    claims, err := auth.ParseToken(tokenString)
    if err != nil {
        respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
        return nil, err
    }

    return claims, nil
}

