package config

import(
	"html/template"
	"net/http"
)


func HandlerTrueBeginnerOne(w http.ResponseWriter, r *http.Request) {
	// Parse the template
	tmpl, err := template.ParseFiles("templates/partials/true_beginner_one.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	// Set the header to ensure correct response type
	w.Header().Set("Content-Type", "text/html")

	// Execute the template without any data
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}