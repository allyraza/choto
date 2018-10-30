package handlers

import (
	"fmt"
	"net/http"

	"strings"

	"github.com/allyraza/choto/repos"
)

// URLHandler :
type URLHandler struct {
	Repo *repos.URLRepo
}

// Index : list all the urls
func (uh *URLHandler) Index(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/")

	if key != "" {
		if url, err := uh.Repo.FindByKey(key); err == nil {
			http.Redirect(w, r, url.Long, http.StatusFound)
		}
	}

	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	return
}

// Create : creates a short url
func (uh *URLHandler) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Short url created successfully.")
}
