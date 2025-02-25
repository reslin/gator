package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/reslin/gator/internal/database"
)

type FeedInfo struct {
	FeedData  database.Feed
	FeedOwner string
}

func handlerAddFeed(s *state, cmd command) error {
	cmdArgs := cmd.Args
	if len(cmdArgs) < 2 {
		return fmt.Errorf("not enough arguments")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	feedName := cmdArgs[0]
	feedURL := cmdArgs[1]

	newFeed, err := s.db.AddFeed(context.Background(), database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      feedName,
		Url:       feedURL,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed added successfully!\n")
	PrintFeed(FeedInfo{
		FeedData:  newFeed,
		FeedOwner: s.cfg.CurrentUserName,
	})
	
	return nil
}

func PrintFeed(feed FeedInfo) {

	fmt.Printf("Name:        %s\n", feed.FeedData.Name)
	fmt.Printf("URL:         %s\n", feed.FeedData.Url)
	fmt.Printf("ID:          %s\n", feed.FeedData.ID)
	fmt.Printf("User:        %s\n", feed.FeedOwner)
	fmt.Printf("Created At:  %s\n", feed.FeedData.CreatedAt)
	fmt.Printf("Updated At:  %s\n", feed.FeedData.UpdatedAt)

}