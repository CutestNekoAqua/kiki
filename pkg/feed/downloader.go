package feed

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"

	"gitea.code-infection.com/efertone/kiki/pkg/model"
)

const (
	AtomXMLNS = "http://www.w3.org/2005/Atom"
)

type FeedHeader struct {
	XMLName xml.Name `xml:"feed"`
	XMLNS   string   `xml:"xmlns,attr"`
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

	if feedHeader.XMLNS == AtomXMLNS {
		return ParseAtom(feed.ID, body), nil
	}

	return nil, errors.New("Unknown feed type")
}
