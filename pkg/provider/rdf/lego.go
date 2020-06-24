package rdf

import "encoding/xml"

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
