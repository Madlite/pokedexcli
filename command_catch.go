package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("No pokemon named to catch")
	}
	pokemonName := args[0]

	res, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at ", pokemonName, "...")

	difficulty := int(float64(res.BaseExperience)*0.1 + 50)
	rng := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)

	if rng > difficulty {
		fmt.Println(pokemonName, " escaped!")
		return nil
	}

	fmt.Println(pokemonName, "was caught!")
	cfg.pokeapiClient.StorePokedex(res)
	return nil
}
