package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func startRepl() {
	mapConfig := initializeMapConfig()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}
		commandString := userInput[0]
		command, ok := getCommands()[commandString]
		if !ok {
			fmt.Printf("Unknown command: %v\n", commandString)
			continue
		}
		err := command.callback(mapConfig)
		if err != nil {
			fmt.Printf("The following error occured: %s\n", err)
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}
