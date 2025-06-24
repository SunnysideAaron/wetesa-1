package server

import (
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
	"web/internal/config"
)

type listClientsTemplateData struct {
	MainMenu string
	Response listClientsAPIResponse
	NextPage int
}

type listClientsAPIResponse struct {
	Clients []client `json:"clients"`
	Page    int      `json:"page"`
	Filters struct {
		Name string `json:"name"`
	} `json:"filters"`
	HasNext bool `json:"hasNext"`
}

type client struct {
	ClientID string `json:"client_id"`
	Name     string `json:"name"`
}

func handleListClients(cfg *config.WebConfig, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// // Default values for pagination
			// page := 0

			// // Parse page from query parameter
			// if pageStr := r.URL.Query().Get("page"); pageStr != "" {
			// 	parsedPage, err := strconv.Atoi(pageStr)
			// 	if err != nil {
			// 		http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			// 		return
			// 	}
			// 	if parsedPage < 0 {
			// 		parsedPage = 0
			// 	}
			// 	page = parsedPage
			// }

			// url := cfg.WebAPIURL + "/clients?page=" + strconv.Itoa(page)

			// name := r.URL.Query().Get("name")

			// if name != "" {
			// 	url += "&name=" + name
			// }

			// println(r.URL.RawQuery)

			// url.Values

			// Other web end points wont be just a pass through. probably?
			url := cfg.WebAPIURL + "/clients?" + r.URL.RawQuery

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

			t := "clients_list"

			data := listClientsTemplateData{
				MainMenu: "Clients",
				Response: responseData,
				// NextPage: page + 1,
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
