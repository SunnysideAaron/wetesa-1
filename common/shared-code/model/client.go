package model

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5/pgtype"
)

// Client represents a client record in the database.
// It contains basic information about a client including their unique identifier,
// name, and optional address.
type Client struct {
	ClientID string      `json:"client_id,omitempty"`
	Name     string      `json:"name,omitempty"`
	Address  pgtype.Text `json:"address,omitempty"`
}

// Valid implements the Validator interface for Client.
// It checks if the required fields are properly set and returns a map of validation errors.
// An empty map indicates the Client is valid.
func (c Client) Valid(_ context.Context) map[string]string {
	problems := make(map[string]string)

	if c.Name == "" {
		problems["name"] = "name is required"
	}

	// Address is optional, so no validation needed

	return problems
}

// LogValue implements slog.LogValuer to provide structured logging support.
// It returns the client's ID as the log value. Keeping sensitive fields out of the log.
func (c Client) LogValue() slog.Value {
	return slog.StringValue(c.ClientID)
}

// ClientFilters contains the available filter options for client queries
type ClientFilters struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

// ClientListResponse represents the response structure for client list requests
type ListClientsAPIResponse struct {
	Success  bool                 `json:"success"`
	Messages []MessageAPIResponse `json:"messages"`
	Previous string               `json:"previous,omitempty"`
	Next     string               `json:"next,omitempty"`
	Clients  []Client             `json:"clients"`
}
