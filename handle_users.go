package main

import (
	"context"
	"fmt"

	"github.com/cmilliron/bootdev-gator-go/internal/database"
)

func handleUsers(s *state, cmd command) error {
	fmt.Printf("\nRunning %s...\n", cmd.Name)

	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching users: \n%w\n", err)
	}
	listUsers(s.cfg.CurrentUserName, users)
	return nil
}

func listUsers(currentUser string, users []database.User) {
	fmt.Println("List of registered users:")
	for _, user := range users {
		output := user.Name
		if currentUser == user.Name {
			output = output + " (current)"
		}
		fmt.Printf(" * %s\n", output)
	}
}