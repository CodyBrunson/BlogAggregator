package main

import (
	"context"
	"fmt"
	"os"
)

func handlerResetDB(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Printf("error deleting all users: %w\n", err)
		os.Exit(1)
	}
	fmt.Println("DB reset successfully!")
	return nil
}
