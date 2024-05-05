package api

import "net/http"

type PokeAPIClient struct {
	BaseURL string
	Client  *http.Client
}

func NewPokeAPIClient() *PokeAPIClient {
	return &PokeAPIClient{
		BaseURL: "https://pokeapi.co/api/v2",
		Client:  &http.Client{},
	}
}
