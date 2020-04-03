package atom

import (
	"encoding/xml"

	"gitea.code-infection.com/efertone/kiki/pkg/feed/cleaner"
	"gitea.code-infection.com/efertone/kiki/pkg/model"
	"jaytaylor.com/html2text"
)

// Feed is the root wrapper
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
}

// Entry represents one entry
type Entry struct {
	ID      string `xml:"id"`
	Title   string `xml:"title"`
	Link    Link   `xml:"link"`
	Content []byte `xml:"content"`
}

// Link represents a <link>
type Link struct {
	XMLName xml.Name `xml:"link"`
	HRef    string   `xml:"href,attr"`
}

// Parse parses the feed
func Parse(feedID uint, body []byte) []*model.Entry {
	var atom Feed
	xml.Unmarshal(body, &atom)

	var entries []*model.Entry = make([]*model.Entry, 0)
	for _, entry := range atom.Entries {
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
			Link:    cleaner.RemoveTrackers(entry.Link.HRef),
			Title:   title,
			Content: text,
		})
	}
	return entries
}
