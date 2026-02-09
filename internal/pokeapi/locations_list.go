package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Madlite/pokedexcli/internal/pokedex"
)

func (c *Client) GetLocations(pageURL *string) (LocationListResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := c.cache.Get(url); ok {
		var cache LocationListResponse
		if err := json.Unmarshal(data, &cache); err != nil {
			return LocationListResponse{}, err
		}
		return cache, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationListResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationListResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationListResponse{}, err
	}
	c.cache.Add(url, data)

	locationsResp := LocationListResponse{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationListResponse{}, err
	}

	return locationsResp, nil
}

func (c *Client) GetExploreArea(area string) (AreaExploreResponse, error) {
	url := baseURL + "/location-area/" + area

	if data, ok := c.cache.Get(url); ok {
		var cache AreaExploreResponse
		if err := json.Unmarshal(data, &cache); err != nil {
			return AreaExploreResponse{}, err
		}
		return cache, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaExploreResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaExploreResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaExploreResponse{}, err
	}
	c.cache.Add(url, data)

	exploreResp := AreaExploreResponse{}
	err = json.Unmarshal(data, &exploreResp)
	if err != nil {
		return AreaExploreResponse{}, err
	}

	return exploreResp, nil
}

func (c *Client) GetPokemon(pokemon string) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + pokemon

	if data, ok := c.cache.Get(url); ok {
		var cache PokemonResponse
		if err := json.Unmarshal(data, &cache); err != nil {
			return PokemonResponse{}, err
		}
		return cache, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}
	c.cache.Add(url, data)
	pokemonResp := PokemonResponse{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return PokemonResponse{}, err
	}

	return pokemonResp, nil
}

func (c *Client) StorePokedex(p PokemonResponse) {
	pokemon := pokedex.Pokemon{
		Name:   p.Name,
		Height: p.Height,
		Weight: p.Weight,
		Stats: pokedex.Stats{
			HP:             p.Stats[0].BaseStat,
			Attack:         p.Stats[1].BaseStat,
			Defense:        p.Stats[2].BaseStat,
			SpecialAttack:  p.Stats[3].BaseStat,
			SpecialDefense: p.Stats[4].BaseStat,
			Speed:          p.Stats[5].BaseStat,
		},
		Types: []string{},
	}

	for _, Type := range p.Types {
		pokemon.Types = append(pokemon.Types, Type.Type.Name)
	}
	c.pokedex.Store(pokemon)
}

func (c *Client) GetPokedex(pokemon string) {
	p, err := c.pokedex.Get(pokemon)
	if err != nil {
		fmt.Println("you have not caught that pokemon")
		return
	}

	fmt.Println("Height:", p.Height)
	fmt.Println("Weight:", p.Weight)
	fmt.Println("Stats:")
	fmt.Println("  -hp:", p.Stats.HP)
	fmt.Println("  -attack:", p.Stats.Attack)
	fmt.Println("  -defense:", p.Stats.Defense)
	fmt.Println("  -special-attack:", p.Stats.SpecialAttack)
	fmt.Println("  -special-defense:", p.Stats.SpecialDefense)
	fmt.Println("  -speed:", p.Stats.Speed)
	fmt.Println("Types:")
	for _, Type := range p.Types {
		fmt.Println("  - ", Type)
	}
}
