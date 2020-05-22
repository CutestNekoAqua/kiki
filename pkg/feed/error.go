package feed

import "fmt"

// URLParseError occurs when the downloader was not able to parse the response body.
type URLParseError struct {
	URL string
}

func (e URLParseError) Error() string {
	return fmt.Sprintf("Unable to parse as xml: %s", e.URL)
}

// UnknownFeedType occurs when downloader meets with an unknown feed format.
type UnknownFeedType struct {
	Type string
}

func (e UnknownFeedType) Error() string {
	return fmt.Sprintf("Unknown feed type: %s", e.Type)
}
