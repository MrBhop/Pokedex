package pokeapi

import (
	"encoding/json"
)

func (p *PokeClient) GetPokemonResponse (pokemonName string) (PokemonResponse, error) {
	var response PokemonResponse

	url := PokemonEndpoint + pokemonName

	content, err := p.get(url)
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(content, &response); err != nil {
		return response, err
	}

	return response, nil
}
