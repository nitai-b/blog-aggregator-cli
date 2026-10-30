package main

import (
	"log"
	"os"

	"github.com/nitai-b/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	appState := state{cfg: &cfg}

	commands := commands{handlers: make(map[string]func(*state, command) error)}
	commands.register("login", handlerLogin)

	arguments := os.Args
	if len(arguments) < 2 {
		log.Fatal("a command is required")
	}

	cmd := command{name: arguments[1], args: arguments[2:]}
	err = commands.run(&appState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
