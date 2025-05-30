package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func initializeMapConfig() *config {
	url := "https://pokeapi.co/api/v2/location-area/"
	return &config{
		Next: &url,
		Previous: nil,
	}
}

func commandMap(config *config) error {
	if config.Next == nil {
		fmt.Println("You are on the last page")
		return nil
	}

	newConfig, err := commandMapHelper(*config.Next)
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

	newConfig, err := commandMapHelper(*config.Previous)
	if err != nil {
		return err
	}

	config.Next = newConfig.Next
	config.Previous = newConfig.Previous
	return nil
}

func commandMapHelper(url string) (config, error) {
	res, err := http.Get(url)
	if err != nil {
		return config{}, err
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return config{}, fmt.Errorf("Status Code: %v", res.StatusCode)
	}

	var mapResponse MapResponse
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&mapResponse); err != nil {
		return config{}, err
	}

	for _, result := range mapResponse.Results {
		fmt.Printf("%v\n", result.Name)
	}


	return config{
		Next: mapResponse.Next,
		Previous: mapResponse.Previous,
	},
	nil
}
