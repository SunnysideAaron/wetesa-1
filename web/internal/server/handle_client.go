package server

import (
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"web/internal/config"
	"web/internal/shared-code/model"

	"github.com/jackc/pgx/v5/pgtype"
)

type clientsGetTemplateData struct {
	MainMenu string
	Request  *http.Request
	Response model.ListClientsAPIResponse
}

func handleClientsGet(cfg *config.WebConfig, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Other web end points wont be just a pass through. probably?
			url := cfg.WebAPIURLInternal + "/clients?" + r.URL.RawQuery

			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			var responseData model.ListClientsAPIResponse
			err = json.Unmarshal(body, &responseData)
			if err != nil {
				log.Fatal(err)
			}

			// Convert API URLs to web client URLs in pagination links
			if responseData.Next != "" {
				responseData.Next = strings.Replace(
					responseData.Next,
					cfg.WebAPIURLExternal,
					cfg.WebURL,
					1,
				)
			}

			if responseData.Previous != "" {
				responseData.Previous = strings.Replace(
					responseData.Previous,
					cfg.WebAPIURLExternal,
					cfg.WebURL,
					1,
				)
			}

			t := "clients_get"

			data := clientsGetTemplateData{
				MainMenu: "Clients",
				Request:  r,
				Response: responseData,
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

type ClientGetTemplateData struct {
	MainMenu string
	//Request  *http.Request
	Response model.GetClientAPIResponse
}

// handleClientRead handles requests to read a specific client
func handleClientGet(cfg *config.WebConfig, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			url := cfg.WebAPIURLInternal + r.URL.Path

			resp, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			var responseData model.GetClientAPIResponse
			err = json.Unmarshal(body, &responseData)
			if err != nil {
				log.Fatal(err)
			}

			t := "client_get"

			data := ClientGetTemplateData{
				MainMenu: "Clients",
				//Request:  r,
				Response: responseData,
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

func handleClientPost(cfg *config.WebConfig, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			url := cfg.WebAPIURLInternal + r.URL.Path
			httpVerb := r.FormValue("HTTPVerb")
			switch httpVerb {
			case "DELETE":
				req, err := http.NewRequest(http.MethodDelete, url, nil)
				if err != nil {
					logger.LogAttrs(
						r.Context(),
						slog.LevelError,
						"error creating DELETE request",
						slog.String("error", err.Error()),
					)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					logger.LogAttrs(
						r.Context(),
						slog.LevelError,
						"error executing DELETE request",
						slog.String("error", err.Error()),
					)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				defer resp.Body.Close()

				// Redirect to clients list after successful deletion
				if resp.StatusCode == http.StatusNoContent {
					http.Redirect(w, r, "/clients", http.StatusSeeOther)
					return
				}

			case "PUT":
				// // TODO: Implement PUT request handling
				// url := cfg.WebAPIURLInternal + r.URL.Path
				// // Read form data and create request body

				// // For now just return not implemented
				// http.Error(w, "PUT not implemented yet", http.StatusNotImplemented)
				return

			default: // POST
				data := model.Client{
					Name: r.FormValue("name"),
					Address: pgtype.Text{
						String: r.FormValue("address"),
						Valid:  true,
					},
				}

				jsonData, err := json.Marshal(data)
				if err != nil {
					logger.LogAttrs(
						r.Context(),
						slog.LevelError,
						"error marshaling client data",
						slog.String("error", err.Error()),
					)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				// Create request to API
				req, err := http.NewRequest(
					http.MethodPost,
					url,
					strings.NewReader(string(jsonData)),
				)
				if err != nil {
					logger.LogAttrs(
						r.Context(),
						slog.LevelError,
						"error creating POST request",
						slog.String("error", err.Error()),
					)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					logger.LogAttrs(
						r.Context(),
						slog.LevelError,
						"error executing POST request",
						slog.String("error", err.Error()),
					)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode == http.StatusCreated || resp.StatusCode == http.StatusOK {
					http.Redirect(w, r, "/clients", http.StatusSeeOther)
					return
				}

				// If we got here, something went wrong with the API call
				body, _ := io.ReadAll(resp.Body)
				logger.LogAttrs(
					r.Context(),
					slog.LevelError,
					"API returned error",
					slog.Int("status", resp.StatusCode),
					slog.String("response", string(body)),
				)
				http.Error(w, "Error creating client", http.StatusInternalServerError)
			}
		},
	)
}

// bla bla
// func handleClientCreate(cfg *config.WebConfig, logger *slog.Logger) http.Handler {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 			println("test")
// 		},
// 	)
// }
