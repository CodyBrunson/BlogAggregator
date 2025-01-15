package main

import (
	"context"
	"fmt"
	"github.com/codybrunson/gator_1/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
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
