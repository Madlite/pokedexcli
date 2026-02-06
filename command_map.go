package main

import (
	"fmt"
	"errors"
)

func commandMap(cfg *config) error {
	res, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = res.Next
	cfg.prevLocationsURL = res.Previous

	for _, result := range res.Results {
		fmt.Println(result.Name)
	}

	return nil
}


func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	res, err := cfg.pokeapiClient.GetLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}
	
	cfg.nextLocationsURL = res.Next
	cfg.prevLocationsURL = res.Previous

	for _, result := range res.Results {
		fmt.Println(result.Name)
	}

	return nil
}