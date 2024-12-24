package config

import (
    "html/template"
    "net/http"
)

func HandlerLoginPage(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/login.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
    
    err = tmpl.Execute(w, nil)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}