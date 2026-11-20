package subscription

import (
	"log/slog"

	"github.com/SychugovKirill/go-user-subscriptions/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc  *usecase.UseCase
	log *slog.Logger
}

func NewHandler(uc *usecase.UseCase, log *slog.Logger) *Handler {
	return &Handler{
		uc:  uc,
		log: log,
	}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	_ = r
	_ = h
}
