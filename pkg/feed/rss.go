package feed

import (
	"encoding/xml"

	"gitea.code-infection.com/efertone/kiki/pkg/model"
	"jaytaylor.com/html2text"
)

type RSSFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title         string    `xml:"title"`
		Link          string    `xml:"link"`
		Description   string    `xml:"description"`
		LastBuildDate string    `xml:"lastBuildDate"`
		Items         []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	ID      string `xml:"guid"`
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
	Content []byte `xml:"description"`
}

func ParseRSS(feedID uint, body []byte) []*model.Entry {
	var rss RSSFeed
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
			Link:    entry.Link,
			Title:   title,
			Content: text,
		})
	}
	return entries
}
