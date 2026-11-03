// Package handlers
package handlers

import (
	"context"

	"github.com/AHMEDxHAGAG/Gator/internal/arguments"
	"github.com/AHMEDxHAGAG/Gator/internal/state"
)

func HandlerReset(s *state.State, cmd arguments.Arguments) error {
	if err := s.DB.ResetUsers(context.Background()); err != nil {
		return err
	}
	if err := s.DB.ResetFeeds(context.Background()); err != nil {
		return err
	}
	if err := s.DB.ResetFeedFollows(context.Background()); err != nil {
		return err
	}
	return nil
}
