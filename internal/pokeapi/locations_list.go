package pokeapi

import (
	"encoding/json"
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
