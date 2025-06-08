package repl

import (
	"fmt"

	"github.com/MrBhop/Pokedex/internal/pokeapi"
)

func commandMap(config *Config, _ ...string) error {
	if config.Next == nil {
		if config.Previous != nil {
			fmt.Println("You are on the last page")
			return nil
		} else {
			firstPageUrl := pokeapi.LocationAreaEndPoint
			config.Next = &firstPageUrl
		}
	}
	return commandMapHelper(*config.Next, config)
}

func commandMapB(config *Config, _ ...string) error {
	if config.Previous == nil {
		fmt.Println("You are on the first page")
		return nil
	}
	return commandMapHelper(*config.Previous, config)
}

func commandMapHelper(url string, config *Config) error {
	response, err := config.Client.GetMapResponse(url)
	if err != nil {
		return err
	}

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}

	config.Next = response.Next
	config.Previous = response.Previous
	return nil
}

