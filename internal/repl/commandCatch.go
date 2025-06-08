package repl

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must specify a pokemon!")
	}

	pokemonName := args[0]
	pokemonData, err := config.Client.GetPokemonResponse(pokemonName)
	if err != nil {
		return err
	}

	baseXp := pokemonData.BaseExperience

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	if pokemonCaughtSuccesfully(baseXp) {
		config.CaughtPokemon[pokemonName] = pokemonData
		fmt.Printf("%v was caught!\n", pokemonName)
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}

	return nil
}

func pokemonCaughtSuccesfully(baseXp int) bool {
	return rand.Intn(baseXp) < 50
}
