package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/SychugovKirill/go-user-subscriptions/api/docs"
	"github.com/SychugovKirill/go-user-subscriptions/internal/app"
	"github.com/SychugovKirill/go-user-subscriptions/internal/config"
	"github.com/joho/godotenv"
)

// @title			Subscriptions API
// @version		1.0
// @description	API for Subscriptions service
// @BasePath		/api
func main() {
	_ = godotenv.Load()
	cfg := config.MustConfig()
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	application := app.NewApp(log)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := application.Start(ctx, cfg); err != nil {
		log.ErrorContext(ctx, "failed to start app", slog.Any("error", err))
	}
	if err := application.Wait(ctx); err != nil {
		log.ErrorContext(ctx, "app error", slog.Any("error", err))
	}
}
