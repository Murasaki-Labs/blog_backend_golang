package app

import (
	"context"
	"github.com/orgs/murasaki-labs/blog-backend/internal/adapters"
	"github.com/orgs/murasaki-labs/blog-backend/internal/adapters/github"
	"github.com/patrickmn/go-cache"
	"log/slog"
	"time"

	"gorm.io/gorm"
)

// Application provides application features (use cases) service.
type Application interface {
	GetContext() context.Context
	GetArticlesList() ([]github.ArticleMeta, error)
	GetArticleBySlug(slug string) ([]byte, error)
}

type Repo interface {
	// GetDb returns the gorm.DB instance we can work with
	GetDb() (*gorm.DB, error)
}

// App implements interface Application.
type App struct {
	ctx     context.Context
	log     *slog.Logger
	clients *adapters.Clients
	cache   *cache.Cache
}

// GetContext returns the context of the app
func (a *App) GetContext() context.Context {
	return a.ctx
}

// New creates and returns new App.
func New(
	c context.Context,
	log *slog.Logger,
	clients *adapters.Clients,
) *App {
	a := &App{
		ctx:     c,
		log:     log,
		clients: clients,
		cache:   cache.New(5*time.Minute, 10*time.Minute),
	}

	return a
}
