package main

import (
	"time"

	"github.com/Madlite/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}

	startRepl(cfg)
}
