package main

import (
	"time"

	"github.com/aumb/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	pokeDex := make(map[string]pokeapi.Pokemon)

	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokeDex,
	}

	startRepl(cfg)
}
