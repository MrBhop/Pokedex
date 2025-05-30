package main

import (
	"os"
	"fmt"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
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
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	help := []string{}
	for _, item := range getCommands() {
		help = append(help, fmt.Sprintf("%v, %v", item.name, item.description))
	}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n%v\n", strings.Join(help, "\n"))
	return nil
}
