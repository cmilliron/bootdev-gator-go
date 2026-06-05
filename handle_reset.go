package main

import (
	"context"
	"fmt"
)

func handleReset(s *state, cmd command) error {
	fmt.Printf("\nRunning %s...\n", cmd.Name)
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error reseting db: \n%w\n", err)
	}

	fmt.Printf("\nDatabase has been successfully reset.\n")
	return nil
}