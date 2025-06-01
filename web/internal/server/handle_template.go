package server

import (
	"log/slog"
	"net/http"
	"web/internal/config"
)

// handleTemplate is for when we don't need to pass any data to the template.
func handleTemplate(cfg *config.WebConfig, logger *slog.Logger, t string) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			rendered, err := renderTemplate(cfg, t, nil)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelError,
					"error rendering template",
					slog.String("template", t),
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte(rendered))
		},
	)
}
