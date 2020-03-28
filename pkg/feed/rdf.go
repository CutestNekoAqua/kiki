package feed

import (
	"encoding/xml"

	"gitea.code-infection.com/efertone/kiki/pkg/model"
	"jaytaylor.com/html2text"
)

type RDFFeed struct {
	XMLName xml.Name
	Items   []RDFItem `xml:"item"`
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		PubDate     string `xml:"pubDate"`
	} `xml:"channel"`
}

type RDFContent struct {
	Encoded []byte `xml:",chardata"`
}

type RDFItem struct {
	ID      string     `xml:"guid"`
	Title   string     `xml:"title"`
	Link    string     `xml:"link"`
	Content RDFContent `xml:"encoded"`
}

func ParseRDF(feedID uint, body []byte) []*model.Entry {
	var rdf RDFFeed
	xml.Unmarshal(body, &rdf)

	var entries []*model.Entry = make([]*model.Entry, 0)

	for _, entry := range rdf.Items {
		title, _ := html2text.FromString(string(entry.Title), html2text.Options{
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
			Link:    RemoveTrackers(entry.Link),
			Title:   title,
			Content: text,
		})
	}
	return entries
}
