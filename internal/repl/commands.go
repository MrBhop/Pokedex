package repl

type cliCommand struct {
	name string
	description string
	callback func(*Config, ...string) error
}

func getCommands() map[string]cliCommand{
	return map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex.",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Displays a help message.",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Displays the next 20 locations.",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Displays the previous 20 locatons.",
			callback: commandMapB,
		},
		"explore": {
			name: "explore <location>",
			description: "Displays a list of the Pokemon at the specified location.",
			callback: commandExplore,
		},
		"catch": {
			name: "catch <pokemon-name>",
			description: "Attempt to catch the specified pokemon.",
			callback: commandCatch,
		},
	}
}
