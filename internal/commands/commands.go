package commands

import (
	"fmt"

	"github.com/AHMEDxHAGAG/Gator/internal/arguments"
	"github.com/AHMEDxHAGAG/Gator/internal/handlers"
	"github.com/AHMEDxHAGAG/Gator/internal/state"
)

type Commands struct {
	Handlers handlers.HandlerMap
}

func (c *Commands) Run(s *state.State, cmd arguments.Arguments) error {
	f, ok := c.Handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("command %s isnt found", cmd.Name)
	}
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
