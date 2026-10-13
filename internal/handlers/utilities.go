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

func printFeed(feed database.Feed) error {
	fmt.Printf("feed: %s\n", feed.Name)
	fmt.Printf("\t- ID: %s\n", feed.ID)
	fmt.Printf("\t- Name: %s\n", feed.Name)
	fmt.Printf("\t- Created By: %s\n", feed.UserID)
	fmt.Printf("\t- Created At: %s\n", feed.CreatedAt.String())
	fmt.Printf("\t- Updated At: %s\n", feed.UpdatedAt.String())
	return nil
}
