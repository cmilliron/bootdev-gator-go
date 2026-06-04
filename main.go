package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cmilliron/bootdev-gator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	gatorState := &state {
		cfg: &cfg,
	}

	commandRegistry := commands{
		registry: map[string]func(*state, command) error{},
	}
	commandRegistry.register("login", handleLogin)

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
}

