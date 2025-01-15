package main

import (
	"context"
	"fmt"
	"github.com/codybrunson/gator_1/internal/database"
	"github.com/google/uuid"
	"os"
	"time"
)

func handlerRegister(s *state, cmd command) error {
	fmt.Printf("Registering...%v\n", cmd.Args[0])
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		fmt.Printf("user already exists: %w\n", err)
		os.Exit(1)
		return err
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Printf("User registered successfully!\n")
	fmt.Printf("%v\n", user)

	return nil
}
