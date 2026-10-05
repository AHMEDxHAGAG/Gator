// Package handlers
package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"
	"uuid"

	"github.com/AHMEDxHAGAG/Gator/internal/arguments"
	"github.com/AHMEDxHAGAG/Gator/internal/database"
	"github.com/AHMEDxHAGAG/Gator/internal/state"
)

type (
	HandlerMap  map[string]HandlerFunc
	HandlerFunc func(*state.State, arguments.Arguments) error
)

func HandlerLogin(s *state.State, cmd arguments.Arguments) error {
	const expected = 1
	if len(cmd.Args) < expected {
		return fmt.Errorf("wrong number of arguments, expected: %d, found: %d", expected, len(cmd.Args))
	}
	userName := strings.Join(cmd.Args, " ")
	user, err := s.DB.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}
	if err := s.Conf.SetUser(user.Name); err != nil {
		return err
	}
	fmt.Printf("user %s has been sent", userName)
	return nil
}

func HandlerRegister(s *state.State, cmd arguments.Arguments) error {
	const expected = 1
	if len(cmd.Args) < expected {
		return fmt.Errorf("wrong number of arguments, expected: %d, found: %d", expected, len(cmd.Args))
	}
	userName := strings.Join(cmd.Args, " ")
	id := uuid.New()
	created_at, updated_at := time.Now(), time.Now()
	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        id.String(),
		Name:      userName,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	})
	if err != nil {
		return err
	}
	if err := s.Conf.SetUser(user.Name); err != nil {
		return err
	}
	return printUser(user)
}

func HandlerReset(s *state.State, cmd arguments.Arguments) error {
	return s.DB.Reset(context.Background())
}

func HandlerGetUsers(s *state.State, cmd arguments.Arguments) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if err := printUser(user); err != nil {
			return err
		}
	}
	return nil
}
