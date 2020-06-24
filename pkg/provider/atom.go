package provider

import (
	"encoding/xml"

	"gitea.code-infection.com/efertone/kiki/pkg/cleaner"
	"gitea.code-infection.com/efertone/kiki/pkg/provider/atom"
	"jaytaylor.com/html2text"
)

const (
	// AtomName is the name of the Atom provider.
	AtomName = "atom"
)

// Atom provider.
type Atom struct {
}

// Name returns with the name of the provider.
func (a Atom) Name() string {
	return AtomName
}

// Parse tries to parse all entries from an Atom []byte content.
func (a *Atom) Parse(body []byte) (entries []*Entry) {
	var atom atom.Feed

	entries = make([]*Entry, 0)

	err := xml.Unmarshal(body, &atom)
	if err != nil {
		return
	}

	for _, entry := range atom.Entries {
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
			Link:    cleaner.RemoveTrackers(entry.Link.HRef),
			Title:   title,
			Content: text,
		})
	}

	return entries
}
