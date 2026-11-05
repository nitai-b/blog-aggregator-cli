package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nitai-b/blog-aggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the login handler expects a single argument, the username")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("\nthe user: '%s' has been set", user.Name)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("the register handler expects a single argument, the username")
	}

	now := time.Now()

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("user: `%s` was created with id: `%s`\n", user.Name, user.ID)
	return nil
}

func handleReset(s *state, cmd command) error {
	err := s.db.DeleteAllUser(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func handleGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		suffix := ""
		if user.Name == s.cfg.CurrentUserName {
			suffix = "(current)"
		}
		fmt.Printf("* %s %s\n", user.Name, suffix)
	}
	return nil
}
