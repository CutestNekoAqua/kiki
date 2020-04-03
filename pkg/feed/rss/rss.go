package rss

import (
	"encoding/xml"

	"gitea.code-infection.com/efertone/kiki/pkg/feed/cleaner"
	"gitea.code-infection.com/efertone/kiki/pkg/model"
	"jaytaylor.com/html2text"
)

// Feed is the root wrapper
type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title         string `xml:"title"`
		Link          string `xml:"link"`
		Description   string `xml:"description"`
		LastBuildDate string `xml:"lastBuildDate"`
		Items         []Item `xml:"item"`
	} `xml:"channel"`
}

// Item represents one entry
type Item struct {
	ID      string `xml:"guid"`
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
	Content []byte `xml:"description"`
}

// Parse parses the feed
func Parse(feedID uint, body []byte) []*model.Entry {
	var rss Feed
	xml.Unmarshal(body, &rss)

	var entries []*model.Entry = make([]*model.Entry, 0)

	for _, entry := range rss.Channel.Items {
		title, _ := html2text.FromString(string(entry.Title), html2text.Options{
			OmitLinks:    true,
			PrettyTables: false,
		})
		text, _ := html2text.FromString(string(entry.Content), html2text.Options{
			OmitLinks:    true,
			PrettyTables: false,
		})

		entries = append(entries, &model.Entry{
			FeedID:  feedID,
			EntryID: entry.ID,
			Link:    cleaner.RemoveTrackers(entry.Link),
			Title:   title,
			Content: text,
		})
	}
	return entries
}
