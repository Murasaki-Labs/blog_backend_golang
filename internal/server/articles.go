package srv

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// handleListArticles serves JSON list of articles
func (s *Server) handleListArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := s.app.GetArticlesList()
	if err != nil {
		s.log.Error("Failed to fetch article list", "error", err)
		http.Error(w, "Unable to load articles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
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
	w.Write(html)
}
