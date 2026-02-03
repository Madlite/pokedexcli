package main

import (
	"fmt"
	"io"
	"log"
)

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area"

	res, err := cfg.pokeapiClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)

	return nil
}