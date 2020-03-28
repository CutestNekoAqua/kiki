package feed

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gitea.code-infection.com/efertone/kiki/pkg/model"
)

type FeedHeader struct {
	XMLName xml.Name
	XMLNS   string `xml:"xmlns,attr"`
}

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

	var feedHeader FeedHeader
	xml.Unmarshal(body, &feedHeader)

	switch strings.ToLower(feedHeader.XMLName.Local) {
	case "feed":
		return ParseAtom(feed.ID, body), nil
	case "rss":
		return ParseRSS(feed.ID, body), nil
	case "rdf":
		return ParseRDF(feed.ID, body), nil
	}

	return nil, errors.New(fmt.Sprintf("Unknown feed type: %s", feedHeader.XMLName.Local))
}
