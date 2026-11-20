package app

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/SychugovKirill/go-user-subscriptions/internal/config"
)

func startHttpServer(_ context.Context, h http.Handler, cfg *config.Http) chan error {
	const (
		readTimeout     = 5 * time.Minute
		writeTimeout    = 5 * time.Minute
		idleTimeout     = 5 * time.Minute
		shutdownTimeout = 1 * time.Minute
	)

	notify := make(chan error, 1)
	server := &http.Server{
		Handler:      h,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
		Addr:         net.JoinHostPort("", strconv.Itoa(int(cfg.Port))),
	}

	go func() {
		notify <- server.ListenAndServe()
		close(notify)
	}()

	return notify
}
