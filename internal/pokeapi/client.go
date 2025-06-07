package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/MrBhop/Pokedex/internal/pokecache"
)

const BaseUrl = "https://pokeapi.co/api/v2/"

type PokeClient struct {
	client http.Client
	cache *pokecache.Cache
}

func NewPokeClient(timeout, cacheInterval time.Duration) PokeClient {
	return PokeClient {
		client: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

func (p *PokeClient) get (url string) ([]byte, error) {
	if item, exists := p.cache.Get(url); exists {
		return item, nil
	}
	res, err := p.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Status Code: %v", res.StatusCode)
	}
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	p.cache.Add(url, content)
	return content, nil
}
