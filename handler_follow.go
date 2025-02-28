package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/reslin/gator/internal/database"
)

func handlerFeedFollow(s *state, cmd command) error {
	cmdArgs := cmd.Args
	if len(cmdArgs) < 1 {
		return fmt.Errorf("not enough arguments")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = FollowFeed(s, cmd.Args[0], user)
	if err != nil {
		return err
	}
	return nil
}

func FollowFeed(s *state, feedURL string, user database.User) error {
	feedFound, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return err
	}

	feedID := feedFound.ID
	createFeedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feedID,
		UserID:    user.ID,
	}

	createdFeedFollow, err := s.db.CreateFeedFollow(context.Background(), createFeedFollow)
	if err != nil {
		return err
	}
	fmt.Printf("Created feed follow: %+v\n", createdFeedFollow)
	return nil
}