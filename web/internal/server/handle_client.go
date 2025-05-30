package server

import (
	"log/slog"
	"net/http"

	"html/template"
)

// Define a struct to hold data for the template
type PageData struct {
	Title   string
	Message string
}

// handleHealthz returns an http.Handler that responds to health check requests
// with a 200 OK status and "OK" message.
func handleListClients(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tmpl, err := template.ParseFiles("template.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Create data for the template
			data := PageData{
				Title:   "My Template",
				Message: "Hello from the template!",
			}

			// Execute the template with the data
			err = tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		},
	)
}
