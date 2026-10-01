// Package handlers
package handlers

import (
	"fmt"
	"strings"

	"github.com/AHMEDxHAGAG/Gator/internal/arguments"
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
	userName := strings.Join(cmd.Args[:], " ")
	if err := s.Conf.SetUser(userName); err != nil {
		return err
	}
	fmt.Printf("user %s has been sent", userName)
	return nil
}
