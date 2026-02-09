package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("No pokemon entered")
	}
	pokemonName := args[0]

	cfg.pokeapiClient.GetPokedex(pokemonName)
	return nil
}
