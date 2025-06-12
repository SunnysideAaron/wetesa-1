package server

import (
	"log/slog"
	"net/http"
	"web/internal/config"
)

type listProductsData struct {
	MainMenu string
	Message  string
}

func handleListProducts(cfg *config.WebConfig, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			t := "products_list"

			data := listProductsData{
				MainMenu: "Products",
				Message:  "Hello from the template!",
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
