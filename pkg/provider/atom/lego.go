package atom

import "encoding/xml"

// Feed is the root wrapper.
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
}

// Entry represents one entry.
type Entry struct {
	ID      string `xml:"id"`
	Title   string `xml:"title"`
	Link    Link   `xml:"link"`
	Content []byte `xml:"content"`
}

// Link represents a <link>.
type Link struct {
	XMLName xml.Name `xml:"link"`
	HRef    string   `xml:"href,attr"`
}
