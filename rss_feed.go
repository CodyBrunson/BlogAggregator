package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "RSSFeed")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	err = xml.Unmarshal(respData, &feed)
	if err != nil {
		return nil, err
	}

	err = cleanHTML(&feed)
	if err != nil {
		return nil, err
	}

	return &feed, nil
}

func cleanHTML(dirtyFeed *RSSFeed) error {
	dirtyFeed.Channel.Title = html.UnescapeString(dirtyFeed.Channel.Title)
	dirtyFeed.Channel.Description = html.UnescapeString(dirtyFeed.Channel.Description)
	for i, item := range dirtyFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		dirtyFeed.Channel.Item[i] = item
	}
	return nil
}
