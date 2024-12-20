package routes

import(
	"net/http"
	"github.com/austinwilson1296/fitted/config"
)

func RegisterRoutes(mux *http.ServeMux, cfg *config.ApiCfg) {
    mux.HandleFunc("POST /api/users", cfg.HandlerCreateUser)
    mux.HandleFunc("GET /api/warmup", cfg.HandlerGenerateWarmUp)
	mux.HandleFunc("POST /api/login", cfg.HandlerLogin)


	mux.HandleFunc("GET /", config.HandlerHomePage)
	mux.HandleFunc("GET /workout", config.HandlerWorkOutPage)
}