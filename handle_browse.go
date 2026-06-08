package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cmilliron/bootdev-gator-go/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.Args) == 1 {
		temp, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err == nil {
			limit = int32(temp)
		}
		fmt.Errorf("Conversion error:", err)
	}

	posts, err := s.db.GetPostForUser(context.Background(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit: limit,
		Offset: 0,
	})
	if err != nil {
		return fmt.Errorf("Could not fetch posts to browse: %s\n", err)
	}

	fmt.Printf("Browsing posts: \n")
	for _, post := range posts {
		fmt.Println(post.Title.String)
		fmt.Printf("\nPosts: %s\n\n", post.Description.String)
	}

	return nil
}


// func browseRepl(limit int) {

// }