// Package server provides HTTP server functionality including request handling,
// middleware, and routing for the API.
package server

import (
	"log/slog"
	"net/http"
	"web/internal/config"
)

// AddRoutes maps all the API routes
// [Map the entire API surface in routes.go](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo)
func AddRoutes(
	//ctx context.Context,
	cfg *config.WebConfig,
	logger *slog.Logger,
	//logLevel *slog.LevelVar,
	//templateCache map[string]*template.Template,
) http.Handler {
	baseMux := http.NewServeMux()

	// Note: HTML is limited to only GET and POST requests.
	// Get should NEVER change state.
	baseMux.Handle(http.MethodGet+" /clients", handleClientsGet(cfg, logger))
	baseMux.Handle(http.MethodGet+" /clients/{id}", handleClientGet(cfg, logger))
	baseMux.Handle(http.MethodPost+" /clients/{id}", handleClientPost(cfg, logger))
	baseMux.Handle(http.MethodPost+" /confirm", handleConfirmPost(cfg, logger))
	baseMux.Handle(http.MethodGet+" /healthz", handleHealthz(logger))
	baseMux.Handle(http.MethodGet+" /home", handleTemplate(cfg, logger, "home"))
	baseMux.Handle(http.MethodGet+" /orders", handleListOrders(cfg, logger))
	baseMux.Handle(http.MethodGet+" /products", handleListProducts(cfg, logger))
	baseMux.Handle(http.MethodGet+" /users", handleListUsers(cfg, logger))

	// due to how go works middleware directly on NotFoundHandler is never called.
	// have to wrap the mux instead.
	baseMux.Handle("/", http.NotFoundHandler())

	// Wrap the entire baseMux with core middleware
	return baseMux
}
