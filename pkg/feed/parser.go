package feed

import (
	"encoding/xml"

	"gitea.code-infection.com/efertone/kiki/pkg/model"
	"jaytaylor.com/html2text"
)

type AtomFeed struct {
	XMLName xml.Name    `xml:"feed"`
	Entries []AtomEntry `xml:"entry"`
}

type AtomEntry struct {
	ID      string `xml:"id"`
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Content []byte `xml:"content"`
}

func ParseAtom(feedID uint, body []byte) []*model.Entry {
	var atom AtomFeed
	xml.Unmarshal(body, &atom)

	var entries []*model.Entry = make([]*model.Entry, 0)
	for _, entry := range atom.Entries {
		text, _ := html2text.FromString(string(entry.Content), html2text.Options{
			PrettyTables: false,
		})

		entries = append(entries, &model.Entry{
			FeedID:  feedID,
			EntryID: entry.ID,
			Link:    entry.Link,
			Title:   entry.Title,
			Content: text,
		})
	}
	return entries
}
