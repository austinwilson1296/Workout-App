package routes

import(
	"net/http"
	"github.com/austinwilson1296/fitted/config"
	"github.com/austinwilson1296/fitted/internal/auth"
)

func RegisterRoutes(mux *http.ServeMux, cfg *config.ApiCfg) {
	//User Endpoints
    mux.HandleFunc("POST /api/users", cfg.HandlerCreateUser)
	mux.HandleFunc("POST /api/login", cfg.HandlerLogin)


	//Exercise Endpoints
	mux.Handle("GET /api/warmup",cfg.HandlerGenerateWarmUp)
	mux.Handle("GET /api/exercise",cfg.HandlerGenerateMainExercise)
	


	// HTML page endpoints
	ServeHTMLPages(mux)
}