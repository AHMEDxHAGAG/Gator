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

func HandlerLogin(s *state.State, cmd arguments.Arguments) error {
	const expected = 1
	if len(cmd.Args) < expected {
		return fmt.Errorf("wrong number of arguments, expected: %d, found: %d", expected, len(cmd.Args))
	}
	userName := strings.Join(cmd.Args, " ")
	user, err := s.DB.GetUserFromName(context.Background(), userName)
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
	createdAt, updatedAt := time.Now(), time.Now()
	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        id.String(),
		Name:      userName,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		return err
	}
	if err := s.Conf.SetUser(user.Name); err != nil {
		return err
	}
	return printUser(user)
}
