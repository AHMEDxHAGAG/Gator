package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AHMEDxHAGAG/Gator/internal/arguments"
	"github.com/AHMEDxHAGAG/Gator/internal/commands"
	"github.com/AHMEDxHAGAG/Gator/internal/config"
	"github.com/AHMEDxHAGAG/Gator/internal/state"
)

func errPanic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	garbageArgOffset = 1
	minNumOfArgs     = 2
)

func main() {
	conf, err := config.Read()
	errPanic(err)
	state := &state.State{
		Conf: &conf,
	}
	commands := commands.NewCommands()
	OSArgs := os.Args
	if len(OSArgs) < minNumOfArgs {
		errPanic(fmt.Errorf("number of expected arguments '%d' is less than the given '%d'", 1, len(OSArgs)-1))
	}
	args := arguments.NewArgument(OSArgs[garbageArgOffset:])
	err = commands.Run(state, args)
	errPanic(err)
}
