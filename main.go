package main

import (
	"github.com/fabian-gubler/pokedexcli/internal/api"
	"github.com/fabian-gubler/pokedexcli/internal/cli"
	"github.com/fabian-gubler/pokedexcli/pkg/config"
)

func main() {

	pokeClient := api.NewPokeAPIClient()

	cfg := &config.Config{
		PokeapiClient: pokeClient,
	}

	cli.RunCLI(cfg)
}
