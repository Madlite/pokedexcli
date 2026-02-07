package pokeapi

import (
	"net/http"
	"time"

	"github.com/Madlite/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient *http.Client
	cache      *pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.DefaultClient,
		cache:      pokecache.NewCache(cacheInterval),
	}
}
