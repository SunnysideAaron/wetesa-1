package server

import (
	"log/slog"
	"net/http"
	"web/internal/config"
)

// Define a struct to hold data for the template
type PageData struct {
	Title   string
	Message string
}

func handleListClients(cfg *config.WebConfig, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Create data for the template
			data := PageData{
				Title:   "My Template",
				Message: "Hello from the template!",
			}

			rendered, err := renderTemplate(cfg, "template.page.tmpl", data)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelError,
					"error rendering template",
					slog.String("template", "template.page.tmpl"),
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte(rendered))
		},
	)
}
