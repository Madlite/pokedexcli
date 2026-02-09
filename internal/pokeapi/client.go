package pokeapi

import (
	"net/http"
	"time"

	"github.com/Madlite/pokedexcli/internal/pokecache"
	"github.com/Madlite/pokedexcli/internal/pokedex"
)

type Client struct {
	httpClient *http.Client
	cache      *pokecache.Cache
	pokedex    *pokedex.Storage
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.DefaultClient,
		cache:      pokecache.NewCache(cacheInterval),
		pokedex:    pokedex.NewPokedex(),
	}
}
