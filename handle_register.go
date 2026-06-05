package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cmilliron/bootdev-gator-go/internal/database"
	"github.com/google/uuid"
)

func handleRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No username provide. Usage: register <username>.\n")
	}

	
	username := cmd.args[0]

	newUserInfo := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: username, 
	}


	
	user, err := s.db.CreateUser(context.Background(), newUserInfo)
	if err != nil {
		log.Fatalf("username %s already exists", username)
	}


	err = s.cfg.SetUser(user.Name)

	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	} 

	fmt.Printf("User %s has created and successfully logged in.\n", user.Name)
	return nil
}