// Package handlers
package handlers

import (
	"github.com/AHMEDxHAGAG/Gator/internal/arguments"
	"github.com/AHMEDxHAGAG/Gator/internal/state"
)

type (
	HandlerMap  map[string]HandlerFunc
	HandlerFunc func(*state.State, arguments.Arguments) error
)
