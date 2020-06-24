package provider

import "fmt"

// UnknownXMLFeedType occurs when downloader meets with an unknown feed format.
type UnknownXMLFeedType struct {
	Type string
}

func (e UnknownXMLFeedType) Error() string {
	return fmt.Sprintf("Unknown feed type: %s", e.Type)
}
