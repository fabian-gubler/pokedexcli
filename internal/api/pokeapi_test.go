package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"github.com/fabian-gubler/pokedexcli/pkg/models"
	"testing"
)

func TestListLocationAreas(t *testing.T) {
	// Mock data that the test server will return
	mockResponse := models.LocationResp{
		Count: 2,
		Next:  nil,
		Previous: nil,
		Results: []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			{Name: "location-1", URL: "https://pokeapi.co/api/v2/location/1"},
			{Name: "location-2", URL: "https://pokeapi.co/api/v2/location/2"},
		},
	}

	// Start a new HTTP test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/location" {
			t.Fatalf("Expected to request '/location', got: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer ts.Close()

	// Initialize a new API client with the test server's URL
	client := PokeAPIClient{
		BaseURL: ts.URL,
		Client:  &http.Client{},
	}

	// Test the ListLocationAreas function
	resp, err := client.ListLocationAreas()
	if err != nil {
		t.Fatalf("ListLocationAreas returned an error: %v", err)
	}

	// Validate the response
	if resp.Count != mockResponse.Count {
		t.Errorf("Expected count %d, got %d", mockResponse.Count, resp.Count)
	}

	if len(resp.Results) != len(mockResponse.Results) {
		t.Errorf("Expected %d results, got %d", len(mockResponse.Results), len(resp.Results))
	}

	for i, result := range resp.Results {
		expected := mockResponse.Results[i]
		if result.Name != expected.Name || result.URL != expected.URL {
			t.Errorf("Expected result %v, got %v", expected, result)
		}
	}
}
