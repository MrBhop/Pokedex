package main

type cliCommand struct {
	name string
	description string
	callback func(config *config) error
}

type config struct {
	Next     *string
	Previous *string
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
