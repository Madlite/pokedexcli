package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("Takes no arguments")
	}

	cfg.pokeapiClient.PrintPokedex()
	return nil
}
