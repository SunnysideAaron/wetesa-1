package server

import (
	"log/slog"
	"net/http"
	"web/internal/config"
)

type ListClientsData struct {
	Message string
}

func handleListClients(cfg *config.WebConfig, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			data := ListClientsData{
				Message: "Hello from the template!",
			}

			rendered, err := renderTemplate(cfg, "clients_list", data)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelError,
					"error rendering template",
					slog.String("template", "clients_list"),
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte(rendered))
		},
	)
}
