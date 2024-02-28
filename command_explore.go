package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]

	detailsResp, err := cfg.pokeapiClient.LocationDetails(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", detailsResp.Name)
	fmt.Println("Found Pokemon:")

	for _, details := range detailsResp.PokemonEncounters {
		fmt.Println("-" + " " + details.Pokemon.Name)
	}
	return nil
}
