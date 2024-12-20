package routes

import(
	"net/http"
	"github.com/austinwilson1296/fitted/config"
)

func RegisterRoutes(mux *http.ServeMux, cfg *config.ApiCfg) {
	//User Endpoints
    mux.HandleFunc("POST /api/users", cfg.HandlerCreateUser)
	mux.HandleFunc("POST /api/login", cfg.HandlerLogin)


	//Exercise Endpoints
	mux.HandleFunc("GET /api/warmup", cfg.HandlerGenerateWarmUp)
	mux.HandleFunc("GET /api/exercise", cfg.HandlerGenerateMainExercise)


	// HTML page endpoints
	mux.HandleFunc("GET /", config.HandlerHomePage)
	mux.HandleFunc("GET /workout", config.HandlerWorkOutPage)
}