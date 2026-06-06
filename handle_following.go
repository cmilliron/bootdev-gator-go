package main

import (
	"context"
	"fmt"

	"github.com/cmilliron/bootdev-gator-go/internal/database"
	// "time"
)



func handleFollowing(s *state, cmd command, user database.User) error {
	fmt.Printf("\nRunning %s...\n", cmd.Name)

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Error fetching feed follows: \n%w\n", err)
	}

	if len(feeds) == 0 {
		fmt.Printf("You aren't following any feeds.")
		return nil
	}

	PrintFollowingFeeds(feeds, user.Name)

	return nil
}


func PrintFollowingFeeds(feeds []database.GetFeedFollowsForUserRow, name string) {
	fmt.Printf("Feeds followed by %s:\n", name)
	for _, feed := range feeds {
		fmt.Printf(" - %s\n", feed.FeedName)
	}
}