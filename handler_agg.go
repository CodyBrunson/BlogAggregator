package main

import (
	"context"
	"fmt"
	"github.com/codybrunson/gator_1/internal/database"
	"github.com/google/uuid"
	"log"
	"strings"
	"time"
)

const timeLayout = "2025-01-15 07:00:50.927689"

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <feed_url>\nExample of acceptable input: 1s, 1m, 1h", cmd.Name)
	}
	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	log.Printf("Collecting feeds every %s...", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) error {
	feedToScrap, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	feedData, err := fetchFeed(context.Background(), feedToScrap.Url)
	if err != nil {
		return err
	}

	feedFetched := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		ID:        feedToScrap.ID,
	}
	err = s.db.MarkFeedFetched(context.Background(), feedFetched)
	if err != nil {
		return err
	}

	err = savePosts(feedData, s, feedToScrap.ID)
	if err != nil {
		return err
	}
	return nil
}

func savePosts(feedData *RSSFeed, s *state, feedID uuid.UUID) error {
	for _, item := range feedData.Channel.Item {
		createPost := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: item.PubDate,
			FeedID:      feedID,
		}
		_, err := s.db.CreatePost(context.Background(), createPost)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			} else {
				return err
			}
		}
	}
	return nil
}
