package server

import (
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"web/internal/config"
)

type listClientsTemplateData struct {
	MainMenu string
	Request  *http.Request
	Response listClientsAPIResponse
}

type MessageAPIResponse struct {
	Severity string `json:"severity"` // should be INFO, WARN, or ERROR
	Message  string `json:"message"`
}

type listClientsAPIResponse struct {
	Success  bool                 `json:"success"`
	Messages []MessageAPIResponse `json:"messages"`
	Previous string               `json:"previous,omitempty"`
	Next     string               `json:"next,omitempty"`
	Clients  []client             `json:"clients"`
}

type client struct {
	ClientID string `json:"client_id"`
	Name     string `json:"name"`
}

func handleListClients(cfg *config.WebConfig, logger *slog.Logger) http.Handler {
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

			var responseData listClientsAPIResponse
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

			t := "clients_list"

			data := listClientsTemplateData{
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
