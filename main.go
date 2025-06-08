package main

import (
	"time"

	"github.com/MrBhop/Pokedex/internal/repl"
)

func main() {
	config := repl.NewConfig(5 * time.Second, 5 * time.Minute)
	repl.StartRepl(config)
}
