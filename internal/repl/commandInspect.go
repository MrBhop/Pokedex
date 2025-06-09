package repl

import (
	"fmt"

	"github.com/MrBhop/Pokedex/internal/pokeapi"
)

func commandInspect(config *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must specify a pokemon")
	}

	pokemonName := args[0]

	pokemon, exists := config.CaughtPokemon[pokemonName]
	if !exists {
		return fmt.Errorf("you haven't caught a %v yet", pokemonName)
	}

	listPokemon(&pokemon)

	return nil
}

func listPokemon(pokemon *pokeapi.PokemonResponse) {
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typeItem := range pokemon.Types {
		fmt.Printf("- %v\n", typeItem.Type.Name)
	}
}
