package commands

import (
	"fmt"

	"github.com/AHMEDxHAGAG/RSSaggregator/internal/arguments"
	"github.com/AHMEDxHAGAG/RSSaggregator/internal/handlers"
	"github.com/AHMEDxHAGAG/RSSaggregator/internal/state"
)

type Commands struct {
	Handlers handlers.HandlerMap
}

func (c *Commands) Run(s *state.State, cmd arguments.Arguments) error {
	f, ok := c.Handlers[cmd.Args[0]]
	if !ok {
		return fmt.Errorf("command %s isnt found", cmd.Args[0])
	}
	cmd.Args = cmd.Args[1:]
	if err := f(s, cmd); err != nil {
		return err
	}
	return nil
}

func (c *Commands) register(name string, f handlers.HandlerFunc) {
	c.Handlers[name] = f
}

func (c *Commands) GetDefaultHandlers() {
	c.register("login", handlers.HandlerLogin)
}

func NewCommands() *Commands {
	cmd := &Commands{}
	cmd.Handlers = make(handlers.HandlerMap)
	cmd.GetDefaultHandlers()
	return cmd
}
