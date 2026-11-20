package app

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"github.com/SychugovKirill/go-user-subscriptions/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

func openDB(ctx context.Context, cfg *config.Postgres, log *slog.Logger) (*bun.DB, error) {
	pgConfig, err := pgx.ParseConfig(cfg.Url)
	if err != nil {
		return nil, err
	}

	sqldb := stdlib.OpenDB(*pgConfig)
	if pingErr := pingDB(ctx, sqldb, log); pingErr != nil {
		return nil, pingErr
	}

	db := bun.NewDB(sqldb, pgdialect.New(),
		maxIdleConnsDB(cfg.MaxIdleConns),
		maxOpenConnsDB(cfg.MaxOpenConns),
		connMaxIdleTimeDB(cfg.ConnMaxIdleTime),
		connMaxLifetimeDB(cfg.ConnMaxLifetime),
	)

	return db, err
}

func pingDB(ctx context.Context, db *sql.DB, log *slog.Logger) error {
	const (
		attempts = 20
		delay    = time.Second * 5
	)

	for i := range attempts {
		err := db.PingContext(ctx)
		if err == nil {
			return nil
		}

		log.Info("failed ping db", slog.Any("error", err))

		if i < attempts-1 {
			select {
			case <-time.After(delay):
				// OK
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	return errors.New("failed to ping db")
}

func maxOpenConnsDB(mx int) bun.DBOption {
	return func(db *bun.DB) {
		db.SetMaxOpenConns(mx)
	}
}

func maxIdleConnsDB(mx int) bun.DBOption {
	return func(db *bun.DB) {
		db.SetMaxIdleConns(mx)
	}
}

func connMaxLifetimeDB(mx time.Duration) bun.DBOption {
	return func(db *bun.DB) {
		db.SetConnMaxLifetime(mx)
	}
}

func connMaxIdleTimeDB(mx time.Duration) bun.DBOption {
	return func(db *bun.DB) {
		db.SetConnMaxIdleTime(mx)
	}
}
