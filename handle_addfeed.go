package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cmilliron/bootdev-gator-go/internal/database"
	"github.com/google/uuid"
)

func handleAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error fetching user: \n%w\n", err)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: feedName,
		Url: feedUrl,
		UserID: user.ID, 
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	
	printFeed(feed)

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Println("Feed created successfully:")
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
	fmt.Println("=====================================")
}