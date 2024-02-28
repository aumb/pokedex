package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {

	fmt.Println("Exploring" + " " + args[0] + "...")

	detailsResp, err := cfg.pokeapiClient.LocationDetails(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, details := range detailsResp.PokemonEncounters {
		fmt.Println("-" + " " + details.Pokemon.Name)
	}
	return nil
}
