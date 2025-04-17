package main

import (
	"context"
	"fmt"
	"github.com/orgs/murasaki-labs/blog-backend/internal/adapters"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/orgs/murasaki-labs/blog-backend/internal/app"
	"github.com/orgs/murasaki-labs/blog-backend/internal/config"
	srv "github.com/orgs/murasaki-labs/blog-backend/internal/server"
)

const contextTimeout = 10 * time.Second

type service struct {
	log *slog.Logger
	cfg *config.Config
	app *app.App
	srv *srv.Server
}

func newService(ctx context.Context, log *slog.Logger) (*service, error) {
	cfg, err := config.FromEnv(ctx)
	if err != nil {
		return nil, err
	}

	client := adapters.MustClients()
	
	application := app.New(ctx, log, client)

	server, err := srv.NewServer(application, cfg, log)
	if err != nil {
		return nil, err
	}

	return &service{
		log: log,
		cfg: cfg,
		app: application,
		srv: server,
	}, nil
}

func (s *service) Serve() error {
	appServer, err := s.srv.Serve()
	if err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	q := <-quit
	s.log.Info(fmt.Sprintf("Receive a signal: %s", q.String()))

	ctx, cancel := context.WithTimeout(context.Background(), contextTimeout)
	defer cancel()

	if err = appServer.Shutdown(ctx); err != nil {
		s.log.Error(fmt.Sprintf("!! Graceful server shutdown failed: %v", err))
		return err
	}

	s.log.Info("Server shut down properly")

	return nil
}
