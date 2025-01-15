package main

import (
	"context"
	"fmt"
	"github.com/codybrunson/gator_1/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	feedURL := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return err
	}
	unfollowParams := database.UnfollowFeedParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}
	err = s.db.UnfollowFeed(context.Background(), unfollowParams)
	if err != nil {
		return err
	}
	fmt.Printf("You have unfollowed %v successfully\n", feedURL)
	return nil
}
