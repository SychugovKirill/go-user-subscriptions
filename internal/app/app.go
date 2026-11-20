package app

import (
	"context"
	"log/slog"

	"github.com/SychugovKirill/go-user-subscriptions/internal/config"
	"github.com/SychugovKirill/go-user-subscriptions/internal/repo/subscription"
	"github.com/SychugovKirill/go-user-subscriptions/internal/usecase"
)

type App struct {
	log            *slog.Logger
	httpServerChan chan error
}

func NewApp(log *slog.Logger) *App {
	return &App{
		log: log,
	}
}

func (a *App) Start(ctx context.Context, cfg *config.Config) error {
	pg, err := openDB(ctx, &cfg.Postgres, a.log)
	if err != nil {
		return err
	}

	uc := usecase.NewUseCase(&usecase.Dependencies{
		Log:              a.log,
		SubscriptionRepo: subscription.NewRepo(pg, a.log),
	})

	r := newRouter(a.log)
	registerRoutes(r.Group("/"), uc, a.log)

	a.httpServerChan = startHttpServer(ctx, r, &cfg.Http)

	return nil
}

func (a *App) Wait(ctx context.Context) error {
	select {
	case err := <-a.httpServerChan:
		a.log.ErrorContext(ctx, "app http server error", slog.String("error", err.Error()))
		return err
	case <-ctx.Done():
	}

	return nil
}
