package main

import (
	"time"

	"github.com/MrBhop/Pokedex/internal/pokeapi"
	"github.com/MrBhop/Pokedex/internal/repl"
)

func main() {
	client := pokeapi.NewPokeClient(5 * time.Second, 5 * time.Minute)
	config := &repl.Config{
		Client: &client,
	}
	repl.StartRepl(config)
}
