package repl

import (
	"fmt"

	"github.com/MrBhop/Pokedex/internal/pokeapi"
)

func commandExplore(config *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must specify a location!")
	}

	location := args[0]
	url := pokeapi.LocationAreaEndPoint + location

	response, err := config.Client.GetLocationAreaResponse(url)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", location)
	fmt.Println("Found Pokemon:")
	for _, encounter := range response.PokemonEncounters {
		fmt.Printf(" - %v\n",encounter.Pokemon.Name)
	}

	return nil
}
