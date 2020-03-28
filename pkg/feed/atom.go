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
	Link    Link   `xml:"link"`
	Content []byte `xml:"content"`
}

type Link struct {
	XMLName xml.Name `xml:"link"`
	HRef    string   `xml:"href,attr"`
}

func ParseAtom(feedID uint, body []byte) []*model.Entry {
	var atom AtomFeed
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
			Link:    entry.Link.HRef,
			Title:   title,
			Content: text,
		})
	}
	return entries
}
