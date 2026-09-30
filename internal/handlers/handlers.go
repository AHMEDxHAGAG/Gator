// Package handlers
package handlers

import (
	"fmt"
	"strings"

	"github.com/AHMEDxHAGAG/RSSaggregator/internal/arguments"
	"github.com/AHMEDxHAGAG/RSSaggregator/internal/state"
)

type (
	HandlerMap  map[string]HandlerFunc
	HandlerFunc func(*state.State, arguments.Arguments) error
)

func HandlerLogin(s *state.State, cmd arguments.Arguments) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("wrong number of arguments, expected: %d, found: %d", 1, len(cmd.Args))
	}
	userName := strings.Join(cmd.Args[:], " ")
	if err := s.Conf.SetUser(userName); err != nil {
		return err
	}
	fmt.Printf("user %s has been sent", userName)
	return nil
}
