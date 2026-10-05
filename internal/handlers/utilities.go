package handlers

import (
	"fmt"

	"github.com/AHMEDxHAGAG/Gator/internal/config"
	"github.com/AHMEDxHAGAG/Gator/internal/database"
)

func printUser(user database.User) error {
	currUserName, err := config.GetCurrUser()
	if err != nil {
		return err
	}
	if currUserName == user.Name {
		fmt.Printf("* %s (current)\n", user.Name)
	} else {
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}
