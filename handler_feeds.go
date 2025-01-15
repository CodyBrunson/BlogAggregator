package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Feeds:")
	fmt.Println()
	for _, feed := range feeds {
		feedOwner, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		PrintFeed(FeedInfo{
			FeedData:  feed,
			FeedOwner: feedOwner.Name,
		})
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("================================")
	return nil
}
