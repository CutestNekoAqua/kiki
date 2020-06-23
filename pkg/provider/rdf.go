package provider

import (
	"encoding/xml"

	"gitea.code-infection.com/efertone/kiki/pkg/cleaner"
	"gitea.code-infection.com/efertone/kiki/pkg/provider/rdf"
	"jaytaylor.com/html2text"
)

// RDF provider.
type RDF struct{}

// Name returns with the name of the provider.
func (a RDF) Name() string {
	return "rdf"
}

// Parse tries to parse all entries from an RDF []byte content.
func (a *RDF) Parse(body []byte) (entries []*Entry) {
	var rdf rdf.Feed

	entries = make([]*Entry, 0)

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

		entries = append(entries, &Entry{
			ID:      entry.ID,
			Link:    cleaner.RemoveTrackers(entry.Link),
			Title:   title,
			Content: text,
		})
	}

	return
}
