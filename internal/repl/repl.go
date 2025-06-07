package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MrBhop/Pokedex/internal/pokeapi"
)

type Config struct {
	Client   *pokeapi.PokeClient
	Next     *string
	Previous *string
}

func StartRepl(config *Config) {
	// Initialize the first map url if the config doesn't specify one.
	if config.Next == nil {
		newNext := pokeapi.BaseUrl + "location-area/"
		config.Next = &newNext
	}
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

		err := command.callback(config)
		if err != nil {
			fmt.Printf("The following error occured: %s\n", err)
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}
