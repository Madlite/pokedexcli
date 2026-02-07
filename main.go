package main

import "github.com/Madlite/pokedexcli/internal/pokeapi"

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(20),
	}

	startRepl(cfg)
}
