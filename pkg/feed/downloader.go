package feed

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"

	"gitea.code-infection.com/efertone/kiki/pkg/feed/atom"
	"gitea.code-infection.com/efertone/kiki/pkg/feed/rdf"
	"gitea.code-infection.com/efertone/kiki/pkg/feed/rss"
	"gitea.code-infection.com/efertone/kiki/pkg/model"
)

// Header is the wrapper object to determine what type of feed we have.
type Header struct {
	XMLName xml.Name
	XMLNS   string `xml:"xmlns,attr"`
}

// Download fetches the content of an URI and returns with a parsed []mode.Entry.
func Download(feed *model.Feed) ([]*model.Entry, error) {
	resp, err := http.Get(feed.URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	var feedHeader Header
	parseError := xml.Unmarshal(body, &feedHeader)

	if parseError != nil {
		return nil, URLParseError{feed.URL}
	}

	switch strings.ToLower(feedHeader.XMLName.Local) {
	case "feed":
		return atom.Parse(feed.ID, body), nil
	case "rss":
		return rss.Parse(feed.ID, body), nil
	case "rdf":
		return rdf.Parse(feed.ID, body), nil
	}

	return nil, UnknownFeedType{feedHeader.XMLName.Local}
}
