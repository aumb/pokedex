package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokeResponse, err := cfg.pokeapiClient.PokemonDetails(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeResponse.Name)

	catchRate := 255 - (pokeResponse.BaseExperience / 4)

	randomNumber := rand.Intn(255)

	if randomNumber < catchRate {
		cfg.pokedex[name] = pokeResponse

		fmt.Printf("%s was caught!\n", pokeResponse.Name)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokeResponse.Name)
	}
	return nil
}
