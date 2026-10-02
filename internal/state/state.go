package state

import (
	"github.com/AHMEDxHAGAG/Gator/internal/config"
	"github.com/AHMEDxHAGAG/Gator/internal/database"
)

type State struct {
	Conf *config.Config
	DB   *database.Queries
}

func NewState(conf *config.Config, db *database.Queries) *State {
	state := &State{
		conf,
		db,
	}
	return state
}
