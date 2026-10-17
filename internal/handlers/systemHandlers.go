// Package handlers
package handlers

import (
	"context"

	"github.com/AHMEDxHAGAG/Gator/internal/arguments"
	"github.com/AHMEDxHAGAG/Gator/internal/state"
)

func HandlerReset(s *state.State, cmd arguments.Arguments) error {
	return s.DB.Reset(context.Background())
}
