package main

import (
	"fmt"
	"time"
)

func handleAgg(s *state, cmd command) error {
	fmt.Printf("\nRunning %s...\n", cmd.Name)
		
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <duration in seconds>", cmd.Name)
	}
	duration := cmd.Args[0]
	scrapeInternval, err := time.ParseDuration(duration)
	if err != nil {
		return fmt.Errorf("error parsing time: %w", err)
	}

	ticker := time.NewTicker(scrapeInternval)

	defer ticker.Stop()

	fmt.Printf("Collecting feeds every %s", duration)
	for t := range ticker.C {
		fmt.Printf("Fetching now: %v\n", t)
		err = scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("couldn't fetch feed: %w", err)
		}
	}

	return nil
}
