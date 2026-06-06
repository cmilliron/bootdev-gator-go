package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/cmilliron/bootdev-gator-go/internal/database"
)

func handleFeeds(s *state, cmd command) error {
	fmt.Printf("\nRunning %s...\n", cmd.Name)

	feeds, err := s.db.GetAllFeedsWithUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching feeds: \n%w\n", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	displayFeeds(feeds)

	return nil
}

func displayFeeds(feeds []database.GetAllFeedsWithUsersRow) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	fmt.Printf("Here are the Feeds: \n")
	fmt.Fprintln(w, "Feed\tUrl\tUser Name")
	for _, feed := range feeds {
		fmt.Fprintf(w, "%s\t%s\t%s\n", feed.Name, feed.Url, feed.Username)
	}
	w.Flush()
}