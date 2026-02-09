package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("No explore area entered")
	}
	areaName := args[0]

	res, err := cfg.pokeapiClient.GetExploreArea(areaName)
	if err != nil {
		return err
	}
	fmt.Println("-", areaName, "-")
	fmt.Println("Exploring", areaName, "...")
	fmt.Println("Found pokemon:")
	for _, pokemon := range res.PokemonEncounters {
		fmt.Println(" - ", pokemon.Pokemon.Name)
	}

	return nil
}
