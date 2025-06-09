package repl

import (
	"fmt"
)

func commandPokedex (config *Config, _ ...string) error {
	if len(*&config.CaughtPokemon) == 0 {
		fmt.Println("You have not caught any pokemon yet.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.CaughtPokemon {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil
}
