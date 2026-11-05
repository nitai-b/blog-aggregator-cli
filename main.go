package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/nitai-b/blog-aggregator/internal/config"
	"github.com/nitai-b/blog-aggregator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	appState := state{cfg: &cfg, db: dbQueries}

	commands := commands{handlers: make(map[string]func(*state, command) error)}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handleReset)
	commands.register("users", handleGetUsers)

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
