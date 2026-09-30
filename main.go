package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AHMEDxHAGAG/RSSaggregator/internal/arguments"
	"github.com/AHMEDxHAGAG/RSSaggregator/internal/commands"
	"github.com/AHMEDxHAGAG/RSSaggregator/internal/config"
	"github.com/AHMEDxHAGAG/RSSaggregator/internal/state"
)

func errPanic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conf, err := config.Read()
	errPanic(err)
	state := &state.State{
		Conf: &conf,
	}
	commands := commands.NewCommands()
	OSArgs := os.Args
	if len(OSArgs) < 2 {
		errPanic(fmt.Errorf("number of expected arguments '%d' is less than the given '%d'", 2, len(OSArgs)))
	}
	args := arguments.NewArgument(OSArgs)
	err = commands.Run(state, args)
	errPanic(err)
}
