package routes

import (
	"net/http"
	"github.com/austinwilson1296/fitted/config"
)

func ServeHTMLPages(mux *http.ServeMux) {
	mux.HandleFunc("/", config.HandlerHomePage)
	mux.HandleFunc("/workout", config.HandlerWorkOutPage)
}
