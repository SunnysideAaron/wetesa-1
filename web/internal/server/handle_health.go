package server

import (
	"log/slog"
	"net/http"
)

// handleHealthz returns an http.Handler that responds to health check requests
// with a 200 OK status and "OK" message.
func handleHealthz(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)

			_, err := w.Write([]byte("OK"))
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"could not write OK response",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		},
	)
}
