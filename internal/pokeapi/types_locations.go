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
	BaseExperience int
	Height         int
	Weight         int
	ID             int
	Name           string
	Stats          []struct {
		BaseStat int
		Effort   int
		Stat     struct {
			Name string
			URL  string
		}
	}
	Types []struct {
		Slot int
		Type struct {
			Name string
			URL  string
		}
	}
}
