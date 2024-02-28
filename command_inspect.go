package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	val, ok := cfg.pokedex[name]

	if ok {

		fmt.Printf("Name: %s\n", val.Name)
		fmt.Printf("Height: %d\n", val.Height)
		fmt.Printf("Weight: %d\n", val.Weight)

		fmt.Println("Stats:")
		for _, details := range val.Stats {
			fmt.Printf("-%s: %d\n", details.Stat.Name, details.BaseStat)
		}

		fmt.Println("Types:")
		for _, details := range val.Types {
			fmt.Printf("- %s\n", details.Type.Name)
		}
	} else {
		fmt.Printf("Pokemon %s is not in your Pokedex\n", name)
	}

	return nil

}
