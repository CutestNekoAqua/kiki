package rss

import "encoding/xml"

// Feed is the root wrapper.
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

// Item represents one entry.
type Item struct {
	ID      string `xml:"guid"`
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
	Content []byte `xml:"description"`
}
