// nolint:dupl
package provider_test

import (
	"testing"

	"gitea.code-infection.com/efertone/kiki/pkg/provider"
	"gitea.code-infection.com/efertone/kiki/test"
)

func TestRDF_Parse_Invalid(t *testing.T) {
	prov := provider.NewProviderByName(provider.RDFName)

	input := []byte(`this is not an xml file`)
	entries := prov.Parse(input)

	if len(entries) > 0 {
		t.Errorf("Expected 0 enties, got %d", len(entries))
	}
}

func TestRDF_Parse_Valid(t *testing.T) {
	prov := provider.NewProviderByName(provider.RDFName)

	input := []byte(test.Must(test.LoadFixture("rdf.xml")))
	entries := prov.Parse(input)

	if len(entries) == 0 {
		t.Errorf("Expected enties, got none of them")
		return
	}

	first := entries[0]

	var expected string

	expected = "https://store.steampowered.com/news/58687/"
	if first.ID != expected {
		t.Errorf("ID expected = %s; got = %s", expected, first.ID)
	}

	expected = "Daily Deal - The Beast Inside, 25% Off"
	if first.Title != expected {
		t.Errorf("Title expected = %s; got = %s", expected, first.Title)
	}

	expected = "https://store.steampowered.com/news/58687/"
	if first.Link != expected {
		t.Errorf("Link expected = %s; got = %s", expected, first.Link)
	}

	expected = "Today's Deal: Save 25% on The Beast Inside"

	excerpt := first.ToModel().Excerpt()[:len(expected)]
	if excerpt != expected {
		t.Errorf("Excerpt expected = %s; got = %s", expected, excerpt)
	}
}
