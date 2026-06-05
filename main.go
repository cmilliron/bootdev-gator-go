package main

import (
	"database/sql"

	"fmt"
	"log"
	"os"

	"github.com/cmilliron/bootdev-gator-go/internal/config"
	"github.com/cmilliron/bootdev-gator-go/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	db, err := sql.Open("postgres", cfg.DbUrl)

	dbQueries := database.New(db)

	gatorState := &state {
		db: dbQueries,
		cfg: &cfg,
	}

	commandRegistry := commands{
		registry: map[string]func(*state, command) error{},
	}
	commandRegistry.register("login", handleLogin)
	commandRegistry.register("register", handleRegister)

	if len(os.Args) < 2 {
		fmt.Printf("Usage: cli <command> [args...]")
		os.Exit(1)
	}
	cmd := command {
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = commandRegistry.run(gatorState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// cfg.SetUser("cody")	
}

type state struct {
	cfg 	*config.Config
	db  	*database.Queries
}

