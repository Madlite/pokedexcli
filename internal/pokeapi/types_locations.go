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