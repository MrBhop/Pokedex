package repl


import (
	"fmt"
	"strings"
)

func commandHelp(config *Config) error {
	help := []string{}
	for _, item := range getCommands() {
		help = append(help, fmt.Sprintf("%v, %v", item.name, item.description))
	}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n%v\n", strings.Join(help, "\n"))
	return nil
}
