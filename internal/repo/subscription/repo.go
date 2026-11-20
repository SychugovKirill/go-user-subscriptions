package subscription

import (
	"log/slog"

	"github.com/uptrace/bun"
)

type Repo struct {
	db  bun.IDB
	log *slog.Logger
}

func NewRepo(db bun.IDB, log *slog.Logger) *Repo {
	return &Repo{db, log}
}
