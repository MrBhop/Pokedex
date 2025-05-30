package main

import (
	"time"

	"github.com/MrBhop/Pokedex/internal/pokecache"
)

type cliCommand struct {
	name string
	description string
	callback func(config *config) error
}

type config struct {
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
}

func InitializeConfig() *config {
	url := "https://pokeapi.co/api/v2/location-area/"
	return &config{
		Next: &url,
		Previous: nil,
		Cache: pokecache.NewCache(5 * time.Second),
	}
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays displays the previous 20 locatons",
			callback: commandMapB,
		},
	}
}
