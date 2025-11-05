package handler

import (
	"net/http"
	"shorten-service/internal/repository"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[1:] // remove leading slash

	repo := repository.URLRepository{}

	originalURL, err := repo.GetOriginalURL(shortKey)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
