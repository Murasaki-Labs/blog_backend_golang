package srv

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// handleListArticles serves JSON list of articles
func (s *Server) handleListArticles(w http.ResponseWriter, _ *http.Request) {
	articles, err := s.app.GetArticlesList()
	if err != nil {
		s.log.Error("Failed to fetch article list", "error", err)
		http.Error(w, "Unable to load articles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(articles)
	if err != nil {
		s.log.Error("Failed to encode article list", "error", err)
		http.Error(w, "Unable to encode articles", http.StatusInternalServerError)
		return
	}
}

// handleGetArticleBySlug fetches article .md and converts to HTML
func (s *Server) handleGetArticleBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	html, err := s.app.GetArticleBySlug(slug)
	if err != nil {
		s.log.Error("Failed to fetch article", "slug", slug, "error", err)
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	_, err = w.Write(html)
	if err != nil {
		s.log.Error("Failed to write articles", "error", err)
		http.Error(w, "Failed to write articles", http.StatusInternalServerError)
	}
}
