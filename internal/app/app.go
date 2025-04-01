package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/tehrelt/test-users-api/internal/config"
	"github.com/tehrelt/test-users-api/internal/transport/http"
)

type App struct {
	cfg  *config.Config
	http *http.Server
}

func newApp(cfg *config.Config, http *http.Server) *App {
	return &App{
		cfg:  cfg,
		http: http,
	}
}

func (a *App) Run(ctx context.Context) error {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, os.Kill)

	go func() {
		sig := <-sigchan
		slog.Info("server shutting down", slog.String("signal", sig.String()))
		a.http.Shutdown(ctx)
	}()

	if err := a.http.Run(); err != nil {
		slog.Error("server failed to start", slog.String("error", err.Error()))
		return err
	}

	return nil
}
