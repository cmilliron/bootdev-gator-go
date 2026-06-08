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
	// fmt.Printf("Read config: %+v\n", cfg)
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
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
	commandRegistry.register("reset", handleReset)
	commandRegistry.register("users", handleUsers)
	commandRegistry.register("agg", handleAgg)
	commandRegistry.register("feeds", handleFeeds)
	
	// Protected Commands
	commandRegistry.register("addfeed", middlewareLoggedIn(handleAddFeed))
	commandRegistry.register("follow", middlewareLoggedIn(handleFollow))
	commandRegistry.register("following", middlewareLoggedIn(handleFollowing))
	commandRegistry.register("unfollow", middlewareLoggedIn(handleUnFollow))
	commandRegistry.register("browse", middlewareLoggedIn(handleBrowse))


	if len(os.Args) < 2 {
		fmt.Printf("Usage: cli <command> [args...]")
		os.Exit(1)
	}
	cmd := command {
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = commandRegistry.run(gatorState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}

type state struct {
	cfg 	*config.Config
	db  	*database.Queries
}

