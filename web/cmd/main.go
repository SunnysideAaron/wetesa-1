// Package main is the entry point for the web service.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"web/internal/config"
	"web/internal/logging"
	"web/internal/server"
)

// run is the actual main function
// [func main() only calls run()](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#func-main-only-calls-run)
func run(
	ctx context.Context,
	cfg *config.WebConfig,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	//logger, logLevel := logging.NewLogger(cfg)
	logger, _ := logging.NewLogger(cfg)
	slog.SetDefault(logger)

	// convert from slog to log for http
	httpLogger := slog.NewLogLogger(logger.Handler(), slog.LevelInfo)

	templateCache := server.NewTemplateCache(ctx, logger, "./templates")

	handle := server.AddRoutes(
		//ctx,
		//cfg,
		logger,
		//logLevel,
		templateCache,
	)

	// Configure the HTTP server
	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.WebHost, cfg.WebPort),
		Handler:      handle,
		ReadTimeout:  cfg.WebReadTimeout,
		WriteTimeout: cfg.WebWriteTimeout,
		IdleTimeout:  cfg.WebIdleTimeout,
		ErrorLog:     httpLogger,
	}

	// Start the server in a goroutine
	serverErrors := make(chan error, 1)
	go func() {
		logger.LogAttrs(
			ctx,
			slog.LevelInfo,
			"server starting",
			slog.Any("config", cfg),
		)
		serverErrors <- httpServer.ListenAndServe()
	}()

	// Wait for interrupt or error
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case <-ctx.Done():
		// [Gracefully shutting down](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#gracefully-shutting-down)
		logger.LogAttrs(
			ctx,
			slog.LevelInfo,
			"shutting down server...",
		)

		// Create shutdown context with timeout
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()

		// Attempt graceful shutdown
		if err := httpServer.Shutdown(shutdownCtx); err != nil { //nolint:contextcheck // Yes we want a new context here.
			return fmt.Errorf("server shutdown error: %w", err)
		}
	}

	return nil
}

func main() {
	ctx := context.Background()
	cfg := config.LoadWebConfig()

	if err := run(ctx, cfg); err != nil {
		slog.LogAttrs(
			ctx,
			slog.LevelError,
			"application error",
			slog.Any("error", err),
		)
		os.Exit(1)
	}
}
