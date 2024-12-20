package routes

import(
	"net/http"
)

func ServeStaticFiles(mux *http.ServeMux) {
	staticFiles := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFiles))
}
