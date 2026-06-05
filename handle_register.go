package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cmilliron/bootdev-gator-go/internal/database"
	"github.com/google/uuid"
)

func handleRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("No username provide. Usage: register <username>.\n")
	}

	username := cmd.Args[0]

	newUserInfo := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: username, 
	}
	
	user, err := s.db.CreateUser(context.Background(), newUserInfo)
	if err != nil {
		return fmt.Errorf("username %s already exists", username)
	}

	err = s.cfg.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	} 
	
	printUser(user)
	fmt.Printf("User %s has created and successfully logged in.\n", user.Name)
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}