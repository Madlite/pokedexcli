package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (c *Client) StorePokedex(pokemon PokemonResponse) {
	fmt.Println("Store", pokemon.Name, "in pokedex...")
	// c.pokedex.StorePokedex(pokemon)
}
