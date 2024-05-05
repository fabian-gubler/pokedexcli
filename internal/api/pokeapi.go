package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fabian-gubler/pokedexcli/pkg/models"
)

type PokeAPIClient struct {
	BaseURL string
	Client  *http.Client
}

func NewPokeAPIClient() *PokeAPIClient {
	return &PokeAPIClient{
		BaseURL: "https://pokeapi.co/api/v2",
		Client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *PokeAPIClient) ListLocationAreas(pageUrl *string) (models.LocationResp, error) {

	endpoint := "/location"
	fullURL := c.BaseURL + endpoint

	if pageUrl != nil{
		fullURL = *pageUrl
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return models.LocationResp{}, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return models.LocationResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return models.LocationResp{}, fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.LocationResp{}, err
	}

	locationResp := models.LocationResp{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return models.LocationResp{}, err
	}

	return locationResp, nil

}
