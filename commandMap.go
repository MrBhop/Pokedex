package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/MrBhop/Pokedex/internal/pokecache"
)

type MapResponse struct {
	Count    int       `json:"count"`
	Next     *string    `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMap(config *config) error {
	if config.Next == nil {
		fmt.Println("You are on the last page")
		return nil
	}
	newConfig, err := commandMapHelper(*config.Next, config.Cache)
	if err != nil {
		return err
	}
	config.Next = newConfig.Next
	config.Previous = newConfig.Previous
	return nil
}

func commandMapB(config *config) error {
	if config.Previous == nil {
		fmt.Println("You are on the first page")
		return nil
	}
	newConfig, err := commandMapHelper(*config.Previous, config.Cache)
	if err != nil {
		return err
	}
	config.Next = newConfig.Next
	config.Previous = newConfig.Previous
	return nil
}

func get(url string, cache *pokecache.Cache) ([]byte, error) {
	if item, exists := cache.Get(url); exists {
		return item, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Status Code: %v", res.StatusCode)
	}
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	cache.Add(url, content)
	return content, nil
}

func commandMapHelper(url string, cache *pokecache.Cache) (config, error) {
	newConfig := config{}
	content, err := get(url, cache)
	if err != nil {
		return newConfig, nil
	}
	var mapResponse MapResponse
	if err := json.Unmarshal(content, &mapResponse); err != nil {
		return newConfig, err
	}
	for _, result := range mapResponse.Results {
		fmt.Printf("%v\n", result.Name)
	}
	newConfig.Next = mapResponse.Next
	newConfig.Previous = mapResponse.Previous
	return newConfig, nil
}
