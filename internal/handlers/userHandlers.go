// Package handlers
package handlers

import (
	"context"

	"github.com/AHMEDxHAGAG/Gator/internal/arguments"
	"github.com/AHMEDxHAGAG/Gator/internal/state"
)

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
