package main

import (
	"context"
	"fmt"
	"github.com/codybrunson/gator_1/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerFeedFollow(s *state, cmd command, user database.User) error {
	err := FollowFeed(s, cmd.Args[0], user)
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
