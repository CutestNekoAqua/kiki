// nolint:dupl
package provider_test

import (
	"testing"

	"gitea.code-infection.com/efertone/kiki/pkg/provider"
	"gitea.code-infection.com/efertone/kiki/test"
)

func TestRSS_Parse_Invalid(t *testing.T) {
	prov := provider.NewProviderByName(provider.RSSName)

	input := []byte(`this is not an xml file`)
	entries := prov.Parse(input)

	if len(entries) > 0 {
		t.Errorf("Expected 0 enties, got %d", len(entries))
	}
}

func TestRSS_Parse_Valid(t *testing.T) {
	prov := provider.NewProviderByName(provider.RSSName)

	input := []byte(test.Must(test.LoadFixture("rss.xml")))
	entries := prov.Parse(input)

	if len(entries) == 0 {
		t.Errorf("Expected enties, got none of them")
		return
	}

	first := entries[0]

	var expected string

	expected = "https://blog.gitea.io/2020/06/gitea-1.12.0-and-1.12.1-are-released/"
	if first.ID != expected {
		t.Errorf("ID expected = %s; got = %s", expected, first.ID)
	}

	expected = "Gitea 1.12.0 and 1.12.1 are released"
	if first.Title != expected {
		t.Errorf("Title expected = %s; got = %s", expected, first.Title)
	}

	expected = "https://blog.gitea.io/2020/06/gitea-1.12.0-and-1.12.1-are-released/"
	if first.Link != expected {
		t.Errorf("Link expected = %s; got = %s", expected, first.Link)
	}

	expected = "We are proud to present the release of Gitea version 1.12.0"

	excerpt := first.ToModel().Excerpt()[:len(expected)]
	if excerpt != expected {
		t.Errorf("Excerpt expected = %s; got = %s", expected, excerpt)
	}
}
