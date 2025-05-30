package server

import (
	"bytes"
	"log/slog"
	"net/http"

	"html/template"
)

// Define a struct to hold data for the template
type PageData struct {
	Title   string
	Message string
}

func handleListClients(logger *slog.Logger, templateCache map[string]*template.Template) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Create data for the template
			data := PageData{
				Title:   "My Template",
				Message: "Hello from the template!",
			}

			useCache := false
			var tc map[string]*template.Template

			if useCache {
				tc = templateCache
			} else {
				tc = NewTemplateCache(r.Context(), logger, "./templates")
			}

			ts, ok := tc["template.page.tmpl"]
			if !ok {
				logger.LogAttrs(
					r.Context(),
					slog.LevelError,
					"template not found in cache",
					slog.String("template", "template.page.tmpl"),
				)
				http.Error(w, "Template not found", http.StatusInternalServerError)
				return
			}

			buf := new(bytes.Buffer)

			// Execute the template with the data
			err := ts.Execute(buf, data)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelError,
					"error executing template",
					slog.String("template", "template.page.tmpl"),
					slog.String("error", err.Error()),
				)

				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			buf.WriteTo(w)

		},
	)
}
