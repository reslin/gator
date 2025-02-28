package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	currentID := user.ID
	allUserFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), currentID)
	if err != nil {
		return err
	}
	fmt.Printf("Following:\n")
	for _, feed := range allUserFeeds {
		fmt.Printf(" - %v\n", feed.FeedName)
	}

	return nil
}