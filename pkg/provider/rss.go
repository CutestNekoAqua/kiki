package provider

import (
	"encoding/xml"

	"gitea.code-infection.com/efertone/kiki/pkg/cleaner"
	"gitea.code-infection.com/efertone/kiki/pkg/provider/rss"
	"jaytaylor.com/html2text"
)

// RSS provider.
type RSS struct{}

// Name returns with the name of the provider.
func (a RSS) Name() string {
	return "rss"
}

// Parse tries to parse all entries from an RSS []byte content.
func (a *RSS) Parse(body []byte) (entries []*Entry) {
	var rss rss.Feed

	entries = make([]*Entry, 0)

	err := xml.Unmarshal(body, &rss)
	if err != nil {
		return
	}

	for _, entry := range rss.Channel.Items {
		title, _ := html2text.FromString(entry.Title, html2text.Options{
			OmitLinks:    true,
			PrettyTables: false,
		})
		text, _ := html2text.FromString(string(entry.Content), html2text.Options{
			OmitLinks:    true,
			PrettyTables: false,
		})

		entries = append(entries, &Entry{
			ID:      entry.ID,
			Link:    cleaner.RemoveTrackers(entry.Link),
			Title:   title,
			Content: text,
		})
	}

	return
}
