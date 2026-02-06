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

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationListResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationListResponse{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationListResponse{}, err
	}

	locationsResp := LocationListResponse{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return LocationListResponse{}, err
	}

	return locationsResp, nil
}