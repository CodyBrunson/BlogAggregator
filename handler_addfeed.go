package main

import (
	"context"
	"fmt"
	"github.com/codybrunson/gator_1/internal/database"
	"github.com/google/uuid"
	"time"
)

type FeedInfo struct {
	FeedData  database.Feed
	FeedOwner string
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	cmdArgs := cmd.Args
	if len(cmdArgs) < 2 {
		return fmt.Errorf("not enough arguments")
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
	fmt.Println()
	fmt.Println("===============================")
	err = FollowFeed(s, feedURL, user)
	if err != nil {
		return err
	}
	return nil
}

func PrintFeed(feed FeedInfo) {

	fmt.Printf("Name:\t\t%s\n", feed.FeedData.Name)
	fmt.Printf("URL:\t\t%s\n", feed.FeedData.Url)
	fmt.Printf("ID:\t\t%s\n", feed.FeedData.ID)
	fmt.Printf("User:\t\t%s\n", feed.FeedOwner)
	fmt.Printf("Created At:\t%s\n", feed.FeedData.CreatedAt)
	fmt.Printf("Updated At:\t%s\n", feed.FeedData.UpdatedAt)

}
