package pokeapi

import (
	"encoding/json"
)

func (p *PokeClient) GetMapResponse (url string) (MapResponse, error) {
	var response MapResponse

	content, err := p.get(url)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(content, &response); err != nil {
		return response, err
	}

	return response, nil
}

func (p *PokeClient) GetLocationAreaResponse (location string) (LocationAreaResponse, error) {
	var response LocationAreaResponse

	url := LocationAreaEndpoint + location

	content, err := p.get(url)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(content, &response); err != nil {
		return response, err
	}

	return response, nil
}
