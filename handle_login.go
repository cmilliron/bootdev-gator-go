package main

import (
	"context"
	"fmt"
	"log"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No username provide. Usage: login <username>.\n")
	}

	username := cmd.args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		log.Fatalf("User %s is not resistered: %w", username, err)
	} 

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	} 

	fmt.Printf("User %s has successfully logged in.\n", user.Name)
	return nil
}