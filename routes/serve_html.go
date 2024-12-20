package routes

import (
	"net/http"
	"github.com/austinwilson1296/fitted/config"
	"github.com/austinwilson1296/fitted/internal/auth"
)

func ServeHTMLPages(mux *http.ServeMux) {
	mux.HandleFunc("/", config.HandlerHomePage)
	mux.HandleFunc("/workout", config.HandlerWorkOutPage)
}
