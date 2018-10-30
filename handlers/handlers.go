package handlers

import (
	"net/http"

	"github.com/allyraza/choto/render"
)

// Index : homepage handler
func Index(w http.ResponseWriter, r *http.Request) {
	// @todo: check setup required

	// @todo: check user is logged in

	render.HTML(w, "index.html", struct{}{})
}
