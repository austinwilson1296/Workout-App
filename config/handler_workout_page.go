package config

import (
	"html/template"
	"log"
	"net/http"
)

// Preload templates during application initialization
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func HandlerWorkOutPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Data to pass to the template (add dynamic data here)
	data := map[string]interface{}{
		"Title": "Workout Page",
	}

	// Execute the template with data
	err := tmpl.ExecuteTemplate(w, "workout.html", data)
	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
