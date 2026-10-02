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
	minNumOfArgs     = 1
)

func main() {
	conf, err := config.Read()
	errPanic(err)
	state := &state.State{
		Conf: &conf,
	}
	commands := commands.NewCommands()
	plainArgs := os.Args[garbageArgOffset:]
	if len(plainArgs) < minNumOfArgs {
		errPanic(fmt.Errorf("number of expected arguments '%d' is less than the given '%d'", 1, len(plainArgs)))
	}
	args := arguments.NewArgument(plainArgs)
	err = commands.Run(state, args)
	errPanic(err)
}
