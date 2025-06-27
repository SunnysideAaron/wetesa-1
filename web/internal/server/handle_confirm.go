package server

import (
	"log/slog"
	"net/http"
	"web/internal/config"
)

type ConfirmTemplateData struct {
	MainMenu string
	Question string
	YesHref  string
	NoHref   string
}

// handleTemplate is for when we don't need to pass any data to the template.
func handleConfirmPost(cfg *config.WebConfig, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			t := "confirm"
			data := ConfirmTemplateData{
				MainMenu: "", // Empty string for MainMenu
				Question: r.FormValue("question"),
				YesHref:  r.FormValue("yes_href"),
				NoHref:   r.FormValue("no_href"),
			}

			rendered, err := renderTemplate(cfg, t, data)
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
