package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/cmilliron/bootdev-gator-go/internal/database"
	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx,"GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error making request: %w", err)
	}
	request.Header.Set("User-Agent", "gator")
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Do(request)
	if err != nil || resp.StatusCode != http.StatusOK {
			return nil, err
		}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, err
	}

	cleanFeed(&feed)

	return &feed, nil
}

func cleanFeed(feed *RSSFeed) {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Item[i] = item
	}

}


func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching next feed: %w", err)
	}

	_, err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("Error updating time: %w\n", err)
	}

	feedItems, err := fetchFeed(context.Background(), nextFeed.Url)

	fmt.Printf("Fetching for %s", nextFeed.Name)
	for _, item := range feedItems.Channel.Item {
		date, err := parseCustomTime(item.PubDate)
		if err != nil {
			fmt.Errorf("couldn't Parce time: %w\n", err)
			date = time.Now()
		}
		fmt.Printf(" - %s: %s\n", item.Title, item.PubDate)
		newPost, err := s.db.CreatePost(
			context.Background(),
			database.CreatePostParams{
				ID: uuid.New(),
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
				Title: sql.NullString {
					String: item.Title,
					Valid: true,
				},
				Url: sql.NullString {
					String: item.Link,
					Valid: true,
				},
				Description: sql.NullString {
					String: item.Description,
					Valid: true,
				},
				PublishedAt: sql.NullTime {
					Time: date,
   					Valid: true,
				},
				Unread: true,
				FeedID: nextFeed.ID,
			},
		)
		if err != nil {
			fmt.Printf("Error creating %s\n %w\n", item.Title, err)
		}
		fmt.Printf("Post %s add to database", newPost.Title.String)		
	}

	return nil
}

func parseCustomTime(dateStr string) (time.Time, error) {
	layout := "Mon, 2 Jan 2006 15:04:05"
	
	return time.Parse(layout, dateStr)
}