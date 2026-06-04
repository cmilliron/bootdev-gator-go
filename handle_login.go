package main

import (
	"fmt"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No username provide. Usage: login <username>.\n")
	}

	username := cmd.args[0]

	err := s.cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	} 

	fmt.Printf("User %s has successfully logged in.\n", username)
	return nil
}