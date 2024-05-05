package config

import (
	"github.com/fabian-gubler/pokedexcli/internal/api"
)

type Config struct {
	NextLocationURL     *string
	PreviousLocationURL *string
	PokeapiClient       *api.PokeAPIClient
}
