package rdf

import (
	"encoding/xml"

	"gitea.code-infection.com/efertone/kiki/pkg/feed/cleaner"
	"gitea.code-infection.com/efertone/kiki/pkg/model"
	"jaytaylor.com/html2text"
)

// Feed is the root wrapper.
type Feed struct {
	XMLName xml.Name
	Items   []Item `xml:"item"`
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		PubDate     string `xml:"pubDate"`
	} `xml:"channel"`
}

// Content is a wrapper for content.
type Content struct {
	Encoded []byte `xml:",chardata"`
}

// Item represents one entry.
type Item struct {
	ID      string  `xml:"guid"`
	Title   string  `xml:"title"`
	Link    string  `xml:"link"`
	Content Content `xml:"encoded"`
}

// Parse parses the feed.
func Parse(feedID uint, body []byte) (entries []*model.Entry) {
	var rdf Feed

	entries = make([]*model.Entry, 0)

	err := xml.Unmarshal(body, &rdf)
	if err != nil {
		return
	}

	for _, entry := range rdf.Items {
		title, _ := html2text.FromString(entry.Title, html2text.Options{
			OmitLinks:    true,
			PrettyTables: false,
		})
		text, _ := html2text.FromString(string(entry.Content.Encoded), html2text.Options{
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

	return
}
