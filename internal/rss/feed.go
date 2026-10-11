// Package rss
package rss

import (
	"context"
	"encoding/xml"
	"html"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	rss := &RSSFeed{}
	if err := xml.NewDecoder(res.Body).Decode(rss); err != nil {
		return nil, err
	}
	unEscapeHTMLStringForRSS(rss)
	return rss, nil
}

func unEscapeHTMLStringForRSS(rss *RSSFeed) {
	rss.Channel.Title = html.UnescapeString(rss.Channel.Title)
	rss.Channel.Description = html.UnescapeString(rss.Channel.Description)
	for i, val := range rss.Channel.Item {
		rss.Channel.Item[i].Title = html.UnescapeString(val.Title)
		rss.Channel.Item[i].Description = html.UnescapeString(val.Description)
	}
}
