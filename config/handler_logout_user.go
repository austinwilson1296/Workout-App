package config

import (
	"net/http"
)

func (cfg *ApiCfg) HandlerLogout(w http.ResponseWriter, r *http.Request) {
	// Attempt to retrieve the cookie
	cookie, err := r.Cookie("jwt-token")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "No cookie present", err)
		return
	}

	// Create a new cookie with the same name to overwrite the existing one
	clearedCookie := &http.Cookie{
		Name:     cookie.Name, // "jwt-token"
		Value:    "",          // Clear the value
		Path:     "/",         // Ensure the same path is used
		HttpOnly: true,
		Secure:   true,
		MaxAge:   -1,          // Invalidate the cookie immediately
		SameSite: http.SameSiteStrictMode,
	}

	// Set the cleared cookie
	http.SetCookie(w, clearedCookie)

	// Redirect to login or other desired endpoint
	w.Header().Set("HX-Redirect", "/login")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logged out successfully"))
}
