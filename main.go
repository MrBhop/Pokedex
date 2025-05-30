package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}
		if userInput[0] == "quit" {
			fmt.Println("Exiting Pokedex ...")
			break
		}

		fmt.Printf("Your command was: %v\n", userInput[0])
	}

}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}
