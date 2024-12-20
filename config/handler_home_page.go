package config

import (
	"log"
	"net/http"
)


func HandlerHomePage(w http.ResponseWriter, r *http.Request) {
	claims, err := auth.ParseAndValidateToken(w, r)
    if err != nil {
        return // If authentication fails, stop the handler execution
    }
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Data to pass to the template (add dynamic data here)
	data := map[string]interface{}{
		"Title": "Workout Page",
	}

	// Execute the template with data
	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
