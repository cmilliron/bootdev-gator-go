package main

import (
	"context"
	"fmt"
)

func handleAgg(s *state, cmd command) error {
	fmt.Printf("\nRunning %s...\n", cmd.Name)
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Println("Here are the posts:")
	// for _, feed := range feeds.Channel.Item {
	// 	title := html.UnescapeString(feed.Title)
	// 	d := html.UnescapeString(feed.Description)
	// 	fmt.Printf(" - %s: %s\n", title, d[:100])
	// }

	fmt.Printf("Feed:\n%v\n", feed)
	return nil
}