//nolint:noctx // context isn't needed for tests
package testintegration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"api/internal/database"
)

func TestClientEndpoints(t *testing.T) {
	serverAddr := getAPIAddress()
	client := &http.Client{}

	// Test creating a new client
	t.Run("Create Client", func(t *testing.T) {
		//nolint:exhaustruct // ClientID is generated by the database
		newClient := database.Client{
			Name: "Test Client",
			Address: struct {
				String string
				Valid  bool
			}{
				String: "123 Test St",
				Valid:  true,
			},
		}

		jsonData, err := json.Marshal(newClient)
		if err != nil {
			t.Fatalf("failed to marshal client data: %v", err)
		}

		resp, err := client.Post(
			serverAddr+"/clients",
			"application/json",
			bytes.NewBuffer(jsonData),
		)
		if err != nil {
			t.Fatalf("failed to create client: %v", err)
		}
		defer func() {
			if err := resp.Body.Close(); err != nil {
				t.Errorf("failed to close response body: %v", err)
			}
		}()

		if resp.StatusCode != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, resp.StatusCode)
		}

		// Read and parse response
		var response map[string]string
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if response["name"] != newClient.Name {
			t.Errorf("expected name %s, got %s", newClient.Name, response["name"])
		}
		if response["address"] != newClient.Address.String {
			t.Errorf("expected address %s, got %s", newClient.Address.String, response["address"])
		}
	})

	// Test listing clients
	t.Run("List Clients", func(t *testing.T) {
		resp, err := client.Get(serverAddr + "/clients")
		if err != nil {
			t.Fatalf("failed to list clients: %v", err)
		}
		defer func() {
			if err = resp.Body.Close(); err != nil {
				t.Errorf("failed to close response body: %v", err)
			}
		}()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}

		// Read and verify response contains expected fields
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}

		var response map[string]any
		if err := json.Unmarshal(body, &response); err != nil {
			t.Fatalf("failed to parse response JSON: %v", err)
		}

		requiredFields := []string{"clients", "page", "size", "sort", "filters", "returned"}
		for _, field := range requiredFields {
			if _, exists := response[field]; !exists {
				t.Errorf("response missing required field: %s", field)
			}
		}
	})

	// Test getting a specific client
	t.Run("Get Client", func(t *testing.T) {
		// First, list clients to get an ID
		resp, err := client.Get(serverAddr + "/clients")
		if err != nil {
			t.Fatalf("failed to list clients: %v", err)
		}

		var listResponse map[string]any
		if err = json.NewDecoder(resp.Body).Decode(&listResponse); err != nil {
			t.Fatalf("failed to decode list response: %v", err)
		}
		if err = resp.Body.Close(); err != nil {
			t.Errorf("failed to close response body: %v", err)
		}

		clients, ok := listResponse["clients"].([]any)
		if !ok {
			t.Fatal("response clients field is not an array")
		}
		if len(clients) == 0 {
			t.Skip("no clients available to test Get endpoint")
		}

		firstClient, ok := clients[0].(map[string]any)
		if !ok {
			t.Fatal("client entry is not a map")
		}
		clientID, ok := firstClient["client_id"].(string)
		if !ok {
			t.Fatal("client_id is not a string")
		}

		// Now get the specific client
		resp, err = client.Get(fmt.Sprintf("%s/clients/%s", serverAddr, clientID))
		if err != nil {
			t.Fatalf("failed to get client: %v", err)
		}
		defer func() {
			if err := resp.Body.Close(); err != nil {
				t.Errorf("failed to close response body: %v", err)
			}
		}()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
	})

	// Test updating a client
	t.Run("Update Client", func(t *testing.T) {
		// First, list clients to get an ID
		resp, err := client.Get(serverAddr + "/clients")
		if err != nil {
			t.Fatalf("failed to list clients: %v", err)
		}

		var listResponse map[string]any
		if err = json.NewDecoder(resp.Body).Decode(&listResponse); err != nil {
			t.Fatalf("failed to decode list response: %v", err)
		}
		if err = resp.Body.Close(); err != nil {
			t.Errorf("failed to close response body: %v", err)
		}

		clients, ok := listResponse["clients"].([]any)
		if !ok {
			t.Fatal("response clients field is not an array")
		}
		if len(clients) == 0 {
			t.Skip("no clients available to test Update endpoint")
		}

		firstClient, ok := clients[0].(map[string]any)
		if !ok {
			t.Fatal("client entry is not a map")
		}
		clientID, ok := firstClient["client_id"].(string)
		if !ok {
			t.Fatal("client_id is not a string")
		}

		//nolint:exhaustruct // ClientID is generated by the database
		updateData := database.Client{
			Name: "Updated Test Client",
			Address: struct {
				String string
				Valid  bool
			}{
				String: "456 Updated St",
				Valid:  true,
			},
		}

		jsonData, err := json.Marshal(updateData)
		if err != nil {
			t.Fatalf("failed to marshal update data: %v", err)
		}

		req, err := http.NewRequest(
			http.MethodPut,
			fmt.Sprintf("%s/clients/%s", serverAddr, clientID),
			bytes.NewBuffer(jsonData),
		)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err = client.Do(req)
		if err != nil {
			t.Fatalf("failed to update client: %v", err)
		}
		defer func() {
			if err := resp.Body.Close(); err != nil {
				t.Errorf("failed to close response body: %v", err)
			}
		}()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
	})

	// Test deleting a client
	t.Run("Delete Client", func(t *testing.T) {
		//nolint:exhaustruct // ClientID is generated by the database
		newClient := database.Client{
			Name: "Client To Delete",
			Address: struct {
				String string
				Valid  bool
			}{
				String: "789 Delete St",
				Valid:  true,
			},
		}

		jsonData, err := json.Marshal(newClient)
		if err != nil {
			t.Fatalf("failed to marshal client data: %v", err)
		}

		resp, err := client.Post(
			serverAddr+"/clients",
			"application/json",
			bytes.NewBuffer(jsonData),
		)
		if err != nil {
			t.Fatalf("failed to create client: %v", err)
		}
		if err = resp.Body.Close(); err != nil {
			t.Errorf("failed to close response body: %v", err)
		}

		// Now list clients to get the ID
		resp, err = client.Get(serverAddr + "/clients")
		if err != nil {
			t.Fatalf("failed to list clients: %v", err)
		}

		var listResponse map[string]any
		if err = json.NewDecoder(resp.Body).Decode(&listResponse); err != nil {
			t.Fatalf("failed to decode list response: %v", err)
		}
		if err = resp.Body.Close(); err != nil {
			t.Errorf("failed to close response body: %v", err)
		}

		clients, ok := listResponse["clients"].([]any)
		if !ok {
			t.Fatal("response clients field is not an array")
		}
		if len(clients) == 0 {
			t.Fatal("no clients available to test Delete endpoint")
		}

		// Find the client we just created
		var clientID string
		for _, c := range clients {
			clientMap, ok := c.(map[string]any)
			if !ok {
				t.Fatal("client entry is not a map")
			}
			name, ok := clientMap["name"].(string)
			if !ok {
				t.Fatal("client name is not a string")
			}
			if name == "Client To Delete" {
				id, ok := clientMap["client_id"].(string)
				if !ok {
					t.Fatal("client_id is not a string")
				}
				clientID = id
				break
			}
		}

		if clientID == "" {
			t.Fatal("could not find client to delete")
		}

		// Delete the client
		req, err := http.NewRequest(
			http.MethodDelete,
			fmt.Sprintf("%s/clients/%s", serverAddr, clientID),
			nil,
		)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		resp, err = client.Do(req)
		if err != nil {
			t.Fatalf("failed to delete client: %v", err)
		}
		defer func() {
			if err := resp.Body.Close(); err != nil {
				t.Errorf("failed to close response body: %v", err)
			}
		}()

		if resp.StatusCode != http.StatusNoContent {
			t.Errorf("expected status code %d, got %d", http.StatusNoContent, resp.StatusCode)
		}
	})
}
