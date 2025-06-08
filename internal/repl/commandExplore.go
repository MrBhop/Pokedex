package repl

import (
	"fmt"
)

func commandExplore(config *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must specify a location!")
	}

	location := args[0]

	response, err := config.Client.GetLocationAreaResponse(location)
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
