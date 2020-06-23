package provider

import (
	"encoding/xml"
	"strings"
)

// XMLHeader is the wrapper object to determine what type of feed we have.
type XMLHeader struct {
	XMLName xml.Name
	XMLNS   string `xml:"xmlns,attr"`
}

// XMLFeedTypeOf tries to make a decision what type of provider
// we should use based on the content.
// If it's not an XML or it's an unknown type,
// then returns with error.
func XMLFeedTypeOf(content []byte) (string, error) {
	var feedHeader XMLHeader
	parseError := xml.Unmarshal(content, &feedHeader)

	if parseError != nil {
		return "", parseError
	}

	switch strings.ToLower(feedHeader.XMLName.Local) {
	case "feed":
		return "atom", nil
	case "rss":
		return "rss", nil
	case "rdf":
		return "rdf", nil
	}

	return "", UnknownXMLFeedType{feedHeader.XMLName.Local}
}
