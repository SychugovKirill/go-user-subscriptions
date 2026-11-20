package app

import (
	"log/slog"
	"net/http"

	"github.com/SychugovKirill/go-user-subscriptions/internal/api/http/subscription"
	"github.com/SychugovKirill/go-user-subscriptions/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func newRouter(log *slog.Logger) *gin.Engine {
	r := gin.New()
	_ = log

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
	})

	r.GET("/healthz", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	return r
}

func registerRoutes(r *gin.RouterGroup, uc *usecase.UseCase, log *slog.Logger) {
	api := r.Group("/api")

	h := subscription.NewHandler(uc, log)
	h.RegisterRoutes(api)
}
