package server

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"

	"api/internal/databasegen"
)

func handleListClients2(logger *slog.Logger, queries *databasegen.Queries) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			page := 0
			size := 10

			// Parse page from query parameter
			if pageStr := r.URL.Query().Get("page"); pageStr != "" {
				parsedPage, err := strconv.Atoi(pageStr)
				if err != nil {
					http.Error(w, "Invalid page parameter", http.StatusBadRequest)
					return
				}
				if parsedPage < 0 {
					parsedPage = 0
				}
				page = parsedPage
			}

			// Parse size from query parameter
			if sizeStr := r.URL.Query().Get("size"); sizeStr != "" {
				parsedSize, err := strconv.Atoi(sizeStr)
				if err != nil {
					http.Error(w, "Invalid size parameter", http.StatusBadRequest)
					return
				}
				// Enforce reasonable size limits
				if parsedSize > 100 {
					parsedSize = 100
				} else if parsedSize < 1 {
					parsedSize = 1
				}
				size = parsedSize
			}

			// Calculate offset from page and size
			offset := page * size

			params := databasegen.ListClientsParams{
				Name: pgtype.Text{String: "%", Valid: true},
				Off:  int32(offset),
				Lim:  int32(size),
			}

			clients, err := queries.ListClients(r.Context(), params)

			response := map[string]any{
				"clients":  clients,
				"page":     1,
				"size":     1,
				"sort":     1,
				"filters":  1,
				"returned": len(clients),
			}

			err = encode(w, r, http.StatusOK, response)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error encoding response",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		},
	)
}
