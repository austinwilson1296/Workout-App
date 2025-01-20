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
    mux.HandleFunc("GET /api/warmup", cfg.AuthMiddleware(cfg.HandlerGenerateWarmUp))
    mux.HandleFunc("GET /api/cooldown", cfg.AuthMiddleware(cfg.HandlerGenerateCoolDown))
    mux.HandleFunc("GET /api/exercise", cfg.AuthMiddleware(cfg.HandlerGenerateMainExercise))


    //HTMX Routes
    mux.HandleFunc("GET /true-beginner-1", cfg.AuthMiddleware(config.HandlerTrueBeginnerOne))


    // HTML Pages
    ServeHTMLPages(mux,cfg)
}

func ServeHTMLPages(mux *http.ServeMux, cfg *config.ApiCfg) {
    mux.HandleFunc("GET /", cfg.AuthMiddleware(config.HandlerHomePage))
    mux.HandleFunc("GET /login", config.HandlerLoginPage) 
    mux.HandleFunc("GET /logout",cfg.HandlerLogout)
    mux.HandleFunc("GET /workout", cfg.AuthMiddleware(config.HandlerWorkOutPage))
}