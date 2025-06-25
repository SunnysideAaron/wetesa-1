package server

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"api/internal/database"
	"api/internal/shared-code/model"
)

// handleListClients handles requests to list all clients
func handleListClients(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			qs, page, err := database.ValidateGetClientsParams(r.Context(), r.URL.Query())

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			clients, hasNext, err := db.GetClients(r.Context(), qs)

			// err = errors.New("test error")

			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error getting clients",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			//TODO make this dynamic or pull from env or something. I'm being lasy atm.
			// r.host?
			baseURL := "http://localhost:8080/api/v0.1"

			previous := ""
			if page > 0 {
				previous = r.URL.String()
				previous = strings.ReplaceAll(previous, "page="+strconv.Itoa(page), "page="+strconv.Itoa(page-1))
				previous = baseURL + previous
			}

			next := ""
			if hasNext {
				next = r.URL.String()
				if strings.Contains(next, "page=") {
					next = strings.ReplaceAll(next, "page="+strconv.Itoa(page), "page="+strconv.Itoa(page+1))
				} else {
					if strings.Contains(next, "?") {
						next += "&page=" + strconv.Itoa(page+1)
					} else {
						next += "?page=" + strconv.Itoa(page+1)
					}
				}
				next = baseURL + next
			}

			response := model.ListClientsAPIResponse{
				Success:  true,
				Previous: previous,
				// Messages
				Next:    next,
				Clients: clients,
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

// handleGetClient handles requests to get a specific client
func handleGetClient(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("id")
			if id == "" {
				http.Error(w, "Missing ID", http.StatusBadRequest)
				return
			}

			client, err := db.GetClient(r.Context(), id)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error getting client",
					slog.String("error", err.Error()),
					slog.String("id", id),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			response := model.GetClientAPIResponse{
				Success: true,
				// Messages
				Client: client,
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

			// example of logging while hiding sensitive information
			logger.LogAttrs(
				r.Context(),
				slog.LevelDebug,
				"client retrieved",
				slog.Any("client", client.LogValue()),
			)
		},
	)
}

// handleCreateClient handles requests to create a new client
func handleCreateClient(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			client, problems, err := decode[model.Client](r)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error decoding request",
					slog.String("error", err.Error()),
					slog.Any("problems", problems),
				)
				if err = encode(w, r, http.StatusBadRequest, map[string]any{
					"error":    err.Error(),
					"problems": problems,
				}); err != nil {
					logger.LogAttrs(
						r.Context(),
						slog.LevelInfo,
						"error encoding response",
						slog.String("error", err.Error()),
					)
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}

			err = db.InsertClient(r.Context(), client)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error creating client",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Return the client data
			response := map[string]string{
				"name":    client.Name,
				"address": client.Address.String,
			}

			err = encode(w, r, http.StatusCreated, response)
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

// handleUpdateClient handles requests to update an existing client
// TODO this is for a PUT request. Which is OK but we might want to use PATCH instead.
func handleUpdateClient(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("id")
			if id == "" {
				http.Error(w, "Missing ID", http.StatusBadRequest)
				return
			}

			// First get the existing client
			_, err := db.GetClient(r.Context(), id)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error getting client",
					slog.String("error", err.Error()),
					slog.String("id", id),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Decode the update request
			updateClient, problems, err := decode[model.Client](r)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error decoding request",
					slog.String("error", err.Error()),
					slog.Any("problems", problems),
				)
				if err = encode(w, r, http.StatusBadRequest, map[string]any{
					"error":    err.Error(),
					"problems": problems,
				}); err != nil {
					logger.LogAttrs(
						r.Context(),
						slog.LevelInfo,
						"error encoding response",
						slog.String("error", err.Error()),
					)
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}

			updateClient.ClientID = id

			// Perform the update
			err = db.UpdateClient(r.Context(), updateClient)
			if err != nil {
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error updating client",
					slog.String("error", err.Error()),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Return the updated client
			response := map[string]string{
				"name":    updateClient.Name,
				"address": updateClient.Address.String,
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

// handleDeleteClient handles requests to delete a client
func handleDeleteClient(logger *slog.Logger, db *database.Postgres) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("id")

			if id == "" {
				http.Error(w, "Missing ID", http.StatusBadRequest)
				return
			}

			err := db.DeleteClient(r.Context(), id)
			if err != nil {
				if strings.Contains(err.Error(), "not found") {
					http.Error(w, "Client not found", http.StatusNotFound)
					return
				}
				logger.LogAttrs(
					r.Context(),
					slog.LevelInfo,
					"error deleting client",
					slog.String("error", err.Error()),
					slog.String("id", id),
				)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Return 204 No Content for successful deletion
			w.WriteHeader(http.StatusNoContent)
		},
	)
}
