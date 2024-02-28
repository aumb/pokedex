package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.pokedex) > 0 {
		fmt.Println("Your Pokedex:")

		for _, details := range cfg.pokedex {
			fmt.Printf("- %s\n", details.Name)
		}
	} else {
		fmt.Println("Your Pokedex is empty try catching some Pokemon!")
	}

	return nil

}
