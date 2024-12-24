package routes

import (
    "net/http"
    "github.com/austinwilson1296/fitted/config"
)

func RegisterRoutes(mux *http.ServeMux, cfg *config.ApiCfg) {
    // Static file serving
    fs := http.FileServer(http.Dir("static"))
    mux.Handle("GET /static/", http.StripPrefix("/static/", fs))
    
    // API Endpoints
    mux.HandleFunc("POST /api/users", cfg.HandlerCreateUser)
    mux.HandleFunc("POST /api/login", cfg.HandlerLogin)
    mux.HandleFunc("GET /api/warmup", cfg.HandlerGenerateWarmUp)
    mux.HandleFunc("GET /api/exercise", cfg.HandlerGenerateMainExercise)

    // HTML Pages
    ServeHTMLPages(mux)
}

func ServeHTMLPages(mux *http.ServeMux) {
    mux.HandleFunc("GET /", config.HandlerHomePage)
    mux.HandleFunc("GET /login", config.HandlerLoginPage) 
    mux.HandleFunc("GET /workout", config.HandlerWorkOutPage)
}