package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	var feed RSSFeed

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return &feed, err
	}

	client := &http.Client{}

	req.Header.Set("User-Agent", "gator")
	resp, err := client.Do(req)
	if err != nil {
		return &feed, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return &feed, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &feed, err
	}

	if err := xml.Unmarshal(body, &feed); err != nil {
		return &feed, err
	}

	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	feed.Channel.Link = html.UnescapeString(feed.Channel.Link)
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)

	for i := range feed.Channel.Item {
		item := &feed.Channel.Item[i]
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		item.Link = html.UnescapeString(item.Link)
		item.PubDate = html.UnescapeString(item.PubDate)
	}

	return &feed, nil
}
