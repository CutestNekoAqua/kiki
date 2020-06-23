package provider

import "fmt"

// XMLParseError occurs when the downloader was not able to parse the response body.
type XMLParseError struct {
	URL string
}

func (e XMLParseError) Error() string {
	return fmt.Sprintf("Unable to parse as xml: %s", e.URL)
}

// UnknownXMLFeedType occurs when downloader meets with an unknown feed format.
type UnknownXMLFeedType struct {
	Type string
}

func (e UnknownXMLFeedType) Error() string {
	return fmt.Sprintf("Unknown feed type: %s", e.Type)
}
