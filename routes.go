package choto

import (
	"net/http"

	"github.com/allyraza/choto/handlers"
	"github.com/allyraza/choto/repos"
)

// Routes : registers url routes
func (app *Choto) Routes() {
	mux := http.NewServeMux()

	urlRepo := &repos.URLRepo{Database: app.Database}
	urlHandler := &handlers.URLHandler{Repo: urlRepo}

	mux.HandleFunc("/home", handlers.Index)
	mux.HandleFunc("/create", urlHandler.Create)
	mux.HandleFunc("/", urlHandler.Index)

	app.Mux = mux
}
