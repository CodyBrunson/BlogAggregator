package main

import (
	"context"
	"fmt"
	"github.com/codybrunson/gator_1/internal/database"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s <number of posts>\nThe number of posts is optional. DEFAULT: 2", cmd.Name)
	}
	var postsToPull int
	var err error
	if len(cmd.Args) == 0 {
		postsToPull = 2
	} else {
		postsToPull, err = strconv.Atoi(cmd.Args[0])
	}
	if err != nil {
		postsToPull = 2
	}

	getPostsForUser := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(postsToPull),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), getPostsForUser)
	if err != nil {
		return err
	}
	fmt.Println("Posts:")
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt, post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
		fmt.Println()
	}

	return nil
}
