package app

import (
	"context"
	"log/slog"

	"gorm.io/gorm"
)

// Application provides application features (use cases) service.
type Application interface {
	// GetContext returns the Context for the application
	GetContext() context.Context
}

type Repo interface {
	// GetDb returns the gorm.DB instance we can work with
	GetDb() (*gorm.DB, error)
}

// App implements interface Application.
type App struct {
	ctx context.Context
	log *slog.Logger
}

// GetContext returns the context of the app
func (a *App) GetContext() context.Context {
	return a.ctx
}

// New creates and returns new App.
func New(
	c context.Context,
	log *slog.Logger,
) *App {
	a := &App{
		ctx: c,
		log: log,
	}

	return a
}
