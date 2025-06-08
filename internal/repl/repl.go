package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MrBhop/Pokedex/internal/pokeapi"
)

type Config struct {
	Client   *pokeapi.PokeClient
	Next     *string
	Previous *string
	CaughtPokemon map[string]pokeapi.PokemonResponse
}

func NewConfig (timeout time.Duration, cacheInterval time.Duration) *Config {
	client := pokeapi.NewPokeClient(timeout, cacheInterval)
	return &Config{
		Client: &client,
		CaughtPokemon: map[string]pokeapi.PokemonResponse{},
	}
}

func StartRepl(config *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		userInput := cleanInput(reader.Text())
		if len(userInput) == 0 {
			continue
		}

		commandString := userInput[0]
		command, ok := getCommands()[commandString]
		if !ok {
			fmt.Printf("Unknown command: %v\n", commandString)
			continue
		}

		args := []string{}
		if len(userInput) > 1 {
			args = userInput[1:]
		}

		err := command.callback(config, args...)
		if err != nil {
			fmt.Printf("The following error occured: %s\n", err)
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}
