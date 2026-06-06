package main

import (
	"context"
	"fmt"

	"github.com/cmilliron/bootdev-gator-go/internal/database"
)

func handleUnFollow(s *state, cmd command, user database.User) error {
	fmt.Printf("\nRunning %s...\n", cmd.Name)
	
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error fetching feeds: \n%w\n", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("Error deleting feed follow: \n%w\n", err)
	}

	fmt.Printf("You unfollowed %s\n", feed.Name)

	return nil
}