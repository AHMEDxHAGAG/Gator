// Package handlers
package handlers

import (
	"context"
	"fmt"
	"time"
	"uuid"

	"github.com/AHMEDxHAGAG/Gator/internal/arguments"
	"github.com/AHMEDxHAGAG/Gator/internal/database"
	"github.com/AHMEDxHAGAG/Gator/internal/rss"
	"github.com/AHMEDxHAGAG/Gator/internal/state"
)

func HandlerAggregate(s *state.State, cmd arguments.Arguments) error {
	rss, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("- %s\n", rss.Channel.Title)
	fmt.Printf("  %s\n", rss.Channel.Description)
	for _, val := range rss.Channel.Item {
		fmt.Printf("\t* %s\n", val.Title)
		fmt.Printf("\t  %s\n", val.Description)
	}
	return nil
}

func HandlerAddFeed(s *state.State, cmd arguments.Arguments) error {
	const expected = 2
	if len(cmd.Args) != expected {
		return fmt.Errorf("wrong number of arguments, expected: %d, found: %d", expected, len(cmd.Args))
	}
	feedName, feedURL := cmd.Args[0], cmd.Args[1]
	id := uuid.New()
	createdAt, updatedAt := time.Now(), time.Now()
	feed, err := s.DB.CreateFeed(context.Background(),
		database.CreateFeedParams{
			ID:        id.String(),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Name:      feedName,
			Url:       feedURL,
			UserID:    s.Conf.CurrentUserName,
		})
	if err != nil {
		return err
	}

	id2 := uuid.New()
	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        id2.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		FeedID:    feed.ID,
		UserID:    s.Conf.CurrentUserName,
	})
	if err != nil {
		return err
	}

	return printFeed(feed)
}

func HandlerListFeeds(s *state.State, cmd arguments.Arguments) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		feeds, err := s.DB.GetFeeds(context.Background(), user.ID)
		if err != nil {
			return err
		}
		for _, feed := range feeds {
			if err := printFeed(feed); err != nil {
				return err
			}
			fmt.Printf("\t- username: %s\n", user.Name)
		}
	}
	return nil
}

func HandlerFollowFeed(s *state.State, cmd arguments.Arguments) error {
	const expected = 1
	if len(cmd.Args) != expected {
		return fmt.Errorf("wrong number of arguments, expected: %d, found: %d", expected, len(cmd.Args))
	}
	feedURL := cmd.Args[0]
	feed, err := s.DB.GetFeedFromURL(context.Background(), feedURL)
	if err != nil {
		return err
	}
	id := uuid.New()
	createdAt, updatedAt := time.Now(), time.Now()
	feeds, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        id.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		FeedID:    feed.ID,
		UserID:    s.Conf.CurrentUserName,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Feed %s was added for user %s", feeds.FeedName, feeds.FeedName)
	return nil
}

func HandlerGetFollowingFeeds(s *state.State, cmd arguments.Arguments) error {
	feedFollows, err := s.DB.GetFeedFollowsForUser(context.Background(), s.Conf.CurrentUserName)
	if err != nil {
		return err
	}
	fmt.Println("Your Followed Feeds:")
	for _, feedFollow := range feedFollows {
		feed, err := s.DB.GetFeedFromID(context.Background(), feedFollow.FeedID)
		if err != nil {
			return err
		}
		fmt.Printf("\t - %s", feed.Name)
	}
	return nil
}
