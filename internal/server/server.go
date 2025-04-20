package srv

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/orgs/murasaki-labs/blog-backend/internal/app"
	"github.com/orgs/murasaki-labs/blog-backend/internal/config"
)

const (
	ReadHeaderTimeout = 5 * time.Second
	ReadTimeout       = 10 * time.Second
	WriteTimeout      = 10 * time.Second
	IdleTimeout       = 30 * time.Second
)

type Server struct {
	app app.Application
	cfg *config.Config
	log *slog.Logger
}

func (s *Server) Serve() (*http.Server, error) {
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Timeout(s.cfg.RequestTimeout))

	r.Group(func(r chi.Router) {
		r.Route("/.well-known", func(r chi.Router) {
			r.Get("/live", liveHandler)
			r.Get("/ready", readyHandler)
		})

		r.Route("/api", func(r chi.Router) {
			r.Get("/articles", s.handleListArticles)
			r.Get("/articles/{slug}", s.handleGetArticleBySlug)
		})
	})

	s.log.Info("Routes:")
	_ = chi.Walk(r, func(method string, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
		s.log.Info("%s %s", method, route)
		return nil
	})

	addr := fmt.Sprintf("%s:%d", s.cfg.BindHost, s.cfg.BindPort)
	srv := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: ReadHeaderTimeout,
		ReadTimeout:       ReadTimeout,
		WriteTimeout:      WriteTimeout,
		IdleTimeout:       IdleTimeout,
	}

	go func() {
		s.log.Info(fmt.Sprintf("Listening on %s", addr))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Error(fmt.Sprintf("!! Listen error: %v", err))
		}
	}()

	return srv, nil
}

func NewServer(app app.Application, cfg *config.Config, log *slog.Logger) (*Server, error) {
	return &Server{
		app: app,
		cfg: cfg,
		log: log,
	}, nil
}
