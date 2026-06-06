package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cmilliron/bootdev-gator-go/internal/database"
	"github.com/google/uuid"
)

func handleFollow(s *state, cmd command, user database.User) error {
	fmt.Printf("\nRunning %s...\n", cmd.Name)
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]


	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error fetching feeds: \n%w\n", err)
	}

	CreateFollow(s, feed.ID, user.ID)

	return nil
}

func CreateFollow(s *state, feedId, userId uuid.UUID) (error) {
	newFollowFeed, err := s.db.CreateFollowFeed(context.Background(), database.CreateFollowFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID: feedId,
		UserID: userId, 
	})
	if err != nil {
		return  fmt.Errorf("Error creating follow feed: \n%w\n", err)
	}

	PrintSingleFollow(newFollowFeed)

	return nil
}

func PrintSingleFollow(row database.CreateFollowFeedRow) {
	const border = "=================================================================="
	
	fmt.Println(border)
	fmt.Println(" 🎉 SUCCESS: NEW FEED FOLLOW CREATED")
	fmt.Println(border)
	
	// Core Relationship Info
	fmt.Printf("  User:        %s \n", row.Username)
	fmt.Printf("  Feed Name:   %s\n", row.FeedName)
	fmt.Printf("  Feed URL:    %s\n", row.FeedUrl)
	
	fmt.Println("------------------------------------------------------------------")
	
	// Database IDs & Metadata (for debugging/logs)
	fmt.Printf("  Follow ID:   %s\n", row.ID)
	fmt.Printf("  User ID:     %s\n", row.UserID)
	fmt.Printf("  Feed ID:     %s\n", row.FeedID)
	fmt.Printf("  Followed At: %s\n", row.CreatedAt.Format("2006-01-02 15:04:05 EST"))
	
	fmt.Println(border)
}