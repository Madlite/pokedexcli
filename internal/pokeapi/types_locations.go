package pokeapi

type LocationListResponse struct {
	Count    int
	Next     *string
	Previous *string
	Results  []Result
}

type Result struct {
	Name string
	URL  string
}

type AreaExploreResponse struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string
			URL  string
		}
	} `json:"pokemon_encounters"`
}

type PokemonResponse struct {
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
